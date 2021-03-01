/*
 * Copyright (C) 2016 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package keybinding

import (
	"github.com/godbus/dbus"
	"github.com/linuxdeepin/go-dbus-factory/com.deepin.daemon.audio"
	"github.com/linuxdeepin/go-dbus-factory/com.deepin.daemon.helper.backlight"
	. "pkg.deepin.io/dde/daemon/keybinding/shortcuts"
	"pkg.deepin.io/gir/gio-2.0"
)

const (
	volumeMin = 0
	volumeMax = 1.5
)

const (
	gsKeyOsdAdjustVolState = "osd-adjust-volume-enabled"
)

type OsdVolumeState int32

// Osd音量调节控制
const (
	VolumeAdjustEnable OsdVolumeState = iota
	VolumeAdjustForbidden
	VolumeAdjustHidden
)

type AudioController struct {
	conn                   *dbus.Conn
	audioDaemon            *audio.Audio
	huaweiMicLedWorkaround *huaweiMicLedWorkaround
	gsKeyboard             *gio.Settings
}

func NewAudioController(sessionConn *dbus.Conn,
	backlightHelper *backlight.Backlight) *AudioController {
	c := &AudioController{
		conn:        sessionConn,
		audioDaemon: audio.NewAudio(sessionConn),
	}
	c.initHuaweiMicLedWorkaround(backlightHelper)
	c.gsKeyboard = gio.NewSettings(gsSchemaKeyboard)
	return c
}

func (c *AudioController) Destroy() {
	if c.huaweiMicLedWorkaround != nil {
		c.huaweiMicLedWorkaround.destroy()
		c.huaweiMicLedWorkaround = nil
	}
}

func (*AudioController) Name() string {
	return "Audio"
}

func (c *AudioController) ExecCmd(cmd ActionCmd) error {
	switch cmd {
	case AudioSinkMuteToggle:
		return c.toggleSinkMute()

	case AudioSinkVolumeUp:
		return c.changeSinkVolume(true)

	case AudioSinkVolumeDown:
		return c.changeSinkVolume(false)

	case AudioSourceMuteToggle:
		return c.toggleSourceMute()

	default:
		return ErrInvalidActionCmd{cmd}
	}
}

func (c *AudioController) toggleSinkMute() error {
	var osd string
	var state = OsdVolumeState(c.gsKeyboard.GetEnum(gsKeyOsdAdjustVolState))

	// 当OsdAdjustVolumeState的值为VolumeAdjustEnable时，才会去执行静音操作
	if VolumeAdjustEnable == state {
		sink, err := c.getDefaultSink()
		if err != nil {
			return err
		}

		mute, err := sink.Mute().Get(0)
		if err != nil {
			return err
		}

		err = sink.SetMute(0, !mute)
		if err != nil {
			return err
		}
		osd = "AudioMute"
	} else if VolumeAdjustForbidden == state {
		osd = "AudioMuteAsh"
	} else {
		return nil
	}

	showOSD(osd)
	return nil
}

func (c *AudioController) toggleSourceMute() error {
	var osd string
	var state = OsdVolumeState(c.gsKeyboard.GetEnum(gsKeyOsdAdjustVolState))

	source, err := c.getDefaultSource()
	if err != nil {
		return err
	}

	mute, err := source.Mute().Get(0)
	if err != nil {
		return err
	}
	mute = !mute

	// 当OsdAdjustVolumeState的值为VolumeAdjustEnable时，才会去执行静音操作
	if VolumeAdjustEnable == state {
		err = source.SetMute(0, mute)
		if err != nil {
			return err
		}

		if mute {
			osd = "AudioMicMuteOn"
		} else {
			osd = "AudioMicMuteOff"
		}
	} else if VolumeAdjustForbidden == state {
		if mute {
			osd = "AudioMicMuteOnAsh"
		} else {
			osd = "AudioMicMuteOffAsh"
		}
	} else {
		return nil
	}

	showOSD(osd)
	return nil
}

func (c *AudioController) changeSinkVolume(raised bool) error {
	var osd string
	var state = OsdVolumeState(c.gsKeyboard.GetEnum(gsKeyOsdAdjustVolState))

	// 当OsdAdjustVolumeState的值为VolumeAdjustEnable时，才会去执行调节音量的操作
	if VolumeAdjustEnable == state {
		sink, err := c.getDefaultSink()
		if err != nil {
			return err
		}

		osd = "AudioUp"
		v, err := sink.Volume().Get(0)
		if err != nil {
			return err
		}

		var step = 0.05
		if !raised {
			step = -step
			osd = "AudioDown"
		}

		maxVolume, err := c.audioDaemon.MaxUIVolume().Get(0)
		if err != nil {
			logger.Warning(err)
			maxVolume = volumeMax
		}

		v += step
		if v < volumeMin {
			v = volumeMin
		} else if v > maxVolume {
			v = maxVolume
		}

		logger.Debug("[changeSinkVolume] will set volume to:", v)
		mute, err := sink.Mute().Get(0)
		if err != nil {
			return err
		}

		if mute {
			err = sink.SetMute(0, false)
			if err != nil {
				logger.Warning(err)
			}
		}

		err = sink.SetVolume(0, v, true)
		if err != nil {
			return err
		}
	} else if VolumeAdjustForbidden == state {
		if raised {
			osd = "AudioUpAsh"
		} else {
			osd = "AudioDownAsh"
		}
	} else {
		return nil
	}

	showOSD(osd)
	return nil
}

func (c *AudioController) getDefaultSink() (*audio.Sink, error) {
	sinkPath, err := c.audioDaemon.DefaultSink().Get(0)
	if err != nil {
		return nil, err
	}

	sink, err := audio.NewSink(c.conn, sinkPath)
	if err != nil {
		return nil, err
	}
	return sink, nil
}

func (c *AudioController) getDefaultSource() (*audio.Source, error) {
	sourcePath, err := c.audioDaemon.DefaultSource().Get(0)
	if err != nil {
		return nil, err
	}
	source, err := audio.NewSource(c.conn, sourcePath)
	if err != nil {
		return nil, err
	}

	return source, nil
}
