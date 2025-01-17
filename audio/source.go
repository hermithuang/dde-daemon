/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
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

package audio

import (
	"fmt"
	"strconv"
	"sync"

	dbus "github.com/godbus/dbus"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/pulse"
)

type Source struct {
	audio       *Audio
	service     *dbusutil.Service
	PropsMu     sync.RWMutex
	index       uint32
	cVolume     pulse.CVolume
	channelMap  pulse.ChannelMap
	Name        string
	Description string
	// 默认的输入音量
	BaseVolume     float64
	Mute           bool
	Volume         float64
	Balance        float64
	SupportBalance bool
	Fade           float64
	SupportFade    bool
	// dbusutil-gen: equal=portsEqual
	Ports      []Port
	ActivePort Port
	// 声卡的索引
	Card uint32
}

func newSource(sourceInfo *pulse.Source, audio *Audio) *Source {
	s := &Source{
		audio:   audio,
		index:   sourceInfo.Index,
		service: audio.service,
	}
	if !isPhysicalDevice(sourceInfo.Name) {
		masterSourceInfo := audio.getSourceInfoByName(sourceInfo.Proplist["device.master_device"])
		if masterSourceInfo == nil {
			logger.Warningf("cannot get master source for %s", sourceInfo.Name)
		} else {
			sourceInfo.Card = masterSourceInfo.Card
			sourceInfo.Ports = masterSourceInfo.Ports
			sourceInfo.ActivePort = masterSourceInfo.ActivePort
			logger.Debugf("create reducing noise source on %s", masterSourceInfo.Name)
		}
	}

	s.update(sourceInfo)
	return s
}

// 如何反馈输入音量？
func (s *Source) SetVolume(value float64, isPlay bool) *dbus.Error {
	logger.Debugf("set source %q volume %f", s.Name, value)
	if !isVolumeValid(value) {
		return dbusutil.ToError(fmt.Errorf("invalid volume value: %v", value))
	}

	if value == 0 {
		value = 0.001
	}
	s.PropsMu.RLock()
	cv := s.cVolume.SetAvg(value)
	s.PropsMu.RUnlock()
	s.audio.context().SetSourceVolumeByIndex(s.index, cv)

	configKeeper.SetVolume(s.audio.getCardNameById(s.Card), s.ActivePort.Name, value)
	err := configKeeper.Save(configKeeperFile)
	if err != nil {
		logger.Warning(err)
		return dbusutil.ToError(err)
	}

	if isPlay {
		playFeedback()
	}
	return nil
}

func (s *Source) SetBalance(value float64, isPlay bool) *dbus.Error {
	if value < -1.00 || value > 1.00 {
		return dbusutil.ToError(fmt.Errorf("invalid volume value: %v", value))
	}

	s.PropsMu.RLock()
	cv := s.cVolume.SetBalance(s.channelMap, value)
	s.PropsMu.RUnlock()
	s.audio.context().SetSourceVolumeByIndex(s.index, cv)

	configKeeper.SetBalance(s.audio.getCardNameById(s.Card), s.ActivePort.Name, value)
	err := configKeeper.Save(configKeeperFile)
	if err != nil {
		logger.Warning(err)
		return dbusutil.ToError(err)
	}

	if isPlay {
		playFeedback()
	}
	return nil
}

func (s *Source) SetFade(value float64) *dbus.Error {
	if value < -1.00 || value > 1.00 {
		return dbusutil.ToError(fmt.Errorf("invalid volume value: %v", value))
	}

	s.PropsMu.RLock()
	cv := s.cVolume.SetFade(s.channelMap, value)
	s.PropsMu.RUnlock()
	s.audio.context().SetSourceVolumeByIndex(s.index, cv)

	playFeedback()
	return nil
}

// 设置同时设置音量和平衡
//
// v: volume音量值
// b: balance左右平衡值
// f: fade前后平衡值
//
func (s *Source) setVBF(v, b, f float64) *dbus.Error {
	if v < -1.00 || v > 1.00 {
		return dbusutil.ToError(fmt.Errorf("invalid volume value: %v", v))
	}

	if b < -1.00 || b > 1.00 {
		return dbusutil.ToError(fmt.Errorf("invalid balance value: %v", b))
	}

	if f < -1.00 || b > 1.00 {
		return dbusutil.ToError(fmt.Errorf("invalid fade value: %v", f))
	}

	s.PropsMu.RLock()
	cv := s.cVolume.SetAvg(v)
	cv = cv.SetBalance(s.channelMap, b)
	cv = cv.SetFade(s.channelMap, f)
	s.PropsMu.RUnlock()
	s.audio.context().SetSourceVolumeByIndex(s.index, cv)

	return nil
}

func (s *Source) SetMute(value bool) *dbus.Error {
	s.audio.context().SetSourceMuteByIndex(s.index, value)

	configKeeper.SetMute(s.audio.getCardNameById(s.Card), s.ActivePort.Name, value)
	err := configKeeper.Save(configKeeperFile)
	if err != nil {
		logger.Warning(err)
		return dbusutil.ToError(err)
	}

	if !value {
		playFeedback()
	}
	return nil
}

func (s *Source) SetPort(name string) *dbus.Error {
	s.audio.context().SetSourcePortByIndex(s.index, name)
	return nil
}

func (s *Source) GetMeter() (meter dbus.ObjectPath, busErr *dbus.Error) {
	id := fmt.Sprintf("source%d", s.index)
	s.audio.mu.Lock()
	m, ok := s.audio.meters[id]
	s.audio.mu.Unlock()
	if ok {
		return m.getPath(), nil
	}

	sourceMeter := pulse.NewSourceMeter(s.audio.ctx, s.index)
	m = newMeter(id, sourceMeter, s.audio)
	meterPath := m.getPath()
	err := s.service.Export(meterPath, m)
	if err != nil {
		return "/", dbusutil.ToError(err)
	}

	s.audio.mu.Lock()
	s.audio.meters[id] = m
	s.audio.mu.Unlock()

	m.core.ConnectChanged(func(v float64) {
		m.PropsMu.Lock()
		m.setPropVolume(v)
		m.PropsMu.Unlock()
	})
	return meterPath, nil
}

func (s *Source) getPath() dbus.ObjectPath {
	return dbus.ObjectPath(dbusPath + "/Source" + strconv.Itoa(int(s.index)))
}

func (*Source) GetInterfaceName() string {
	return dbusInterface + ".Source"
}

func (s *Source) update(sourceInfo *pulse.Source) {
	s.PropsMu.Lock()
	s.cVolume = sourceInfo.Volume
	s.channelMap = sourceInfo.ChannelMap
	s.Name = sourceInfo.Name
	s.Description = sourceInfo.Description
	s.Card = sourceInfo.Card
	s.BaseVolume = sourceInfo.BaseVolume.ToPercent()

	s.setPropVolume(floatPrecision(sourceInfo.Volume.Avg()))
	s.setPropMute(sourceInfo.Mute)

	//TODO: handle this
	s.setPropSupportFade(false)
	s.setPropFade(sourceInfo.Volume.Fade(sourceInfo.ChannelMap))
	s.setPropSupportBalance(true)
	s.setPropBalance(sourceInfo.Volume.Balance(sourceInfo.ChannelMap))

	var ports []Port
	for _, p := range sourceInfo.Ports {
		ports = append(ports, toPort(p))
	}

	if isBluezAudio(s.Name) {
		logger.Debugf("create bluez virtual port for source %s", s.Name)
		s.setPropPorts(createBluezVirtualSourcePorts(ports))
		activePort := toPort(sourceInfo.ActivePort)
		activePort.Name += "(headset_head_unit)"
		activePort.Description += "(Headset)"
		s.setPropActivePort(activePort)
	} else {
		s.setPropPorts(ports)
		s.setPropActivePort(toPort(sourceInfo.ActivePort))
	}

	s.PropsMu.Unlock()
}

func (s *Source) setMute(v bool) {
	logger.Debugf("Source #%d setMute %v", s.index, v)
	s.audio.context().SetSourceMuteByIndex(s.index, v)
}
