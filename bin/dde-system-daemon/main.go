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

package main

import (
	"os"

	// modules:
	_ "pkg.deepin.io/dde/daemon/accounts"
	_ "pkg.deepin.io/dde/daemon/apps"
	_ "pkg.deepin.io/dde/daemon/fprintd"
	_ "pkg.deepin.io/dde/daemon/image_effect"
	_ "pkg.deepin.io/dde/daemon/system/airplane_mode"
	_ "pkg.deepin.io/dde/daemon/system/gesture"
	_ "pkg.deepin.io/dde/daemon/system/inputdevices"
	_ "pkg.deepin.io/dde/daemon/system/keyevent"
	_ "pkg.deepin.io/dde/daemon/system/network"
	_ "pkg.deepin.io/dde/daemon/system/power"
	_ "pkg.deepin.io/dde/daemon/system/power_manager"
	_ "pkg.deepin.io/dde/daemon/system/swapsched"
	_ "pkg.deepin.io/dde/daemon/system/systeminfo"
	_ "pkg.deepin.io/dde/daemon/system/timedated"
	_ "pkg.deepin.io/dde/daemon/system/uadp"

	login1 "github.com/linuxdeepin/go-dbus-factory/org.freedesktop.login1"
	"pkg.deepin.io/dde/daemon/loader"
	"pkg.deepin.io/gir/glib-2.0"
	"pkg.deepin.io/lib/dbusutil"
	. "pkg.deepin.io/lib/gettext"
	"pkg.deepin.io/lib/keyfile"
	"pkg.deepin.io/lib/log"
)

//go:generate dbusutil-gen em -type Daemon

type Daemon struct {
	loginManager  *login1.Manager
	systemSigLoop *dbusutil.SignalLoop
	service       *dbusutil.Service
	signals       *struct { //nolint
		HandleForSleep struct {
			start bool
		}
	}
}

const (
	dbusServiceName = "com.deepin.daemon.Daemon"
	dbusPath        = "/com/deepin/daemon/Daemon"
	dbusInterface   = dbusServiceName
)

var logger = log.NewLogger("daemon/dde-system-daemon")
var _daemon *Daemon

func main() {
	service, err := dbusutil.NewSystemService()
	if err != nil {
		logger.Fatal("failed to new system service", err)
	}

	hasOwner, err := service.NameHasOwner(dbusServiceName)
	if err != nil {
		logger.Fatal("failed to call NameHasOwner:", err)
	}
	if hasOwner {
		logger.Warningf("name %q already has the owner", dbusServiceName)
		os.Exit(1)
	}

	// fix no PATH when was launched by dbus
	if os.Getenv("PATH") == "" {
		logger.Warning("No PATH found, manual special")
		err = os.Setenv("PATH", "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin")
		if err != nil {
			logger.Warning(err)
		}
	}

	// 系统级服务，无需设置LANG和LANGUAGE，保证翻译不受到影响
	_ = os.Setenv("LANG", "")
	_ = os.Setenv("LANGUAGE", "")

	InitI18n()
	Textdomain("dde-daemon")

	logger.SetRestartCommand("/usr/lib/deepin-daemon/dde-system-daemon")

	_daemon = &Daemon{}
	_daemon.service = service
	err = service.Export(dbusPath, _daemon)
	if err != nil {
		logger.Fatal("failed to export:", err)
	}

	err = service.RequestName(dbusServiceName)
	if err != nil {
		logger.Fatal("failed to request name:", err)
	}

	startBacklightHelperAsync(service.Conn())
	loader.SetService(service)
	loader.StartAll()
	defer loader.StopAll()

	// NOTE: system/power module requires glib loop
	go glib.StartLoop()

	fixDeepinInstallConfig()
	err = _daemon.forwardPrepareForSleepSignal(service)
	if err != nil {
		logger.Warning(err)
	}
	service.Wait()
}

func (*Daemon) GetInterfaceName() string {
	return dbusInterface
}

func fixDeepinInstallConfig() {
	const (
		filename = "/etc/deepin-installer.conf"
		section  = "General"
		key      = "DI_PASSWORD"
	)

	kf := keyfile.NewKeyFile()
	err := kf.LoadFromFile(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			logger.Warning("failed to load:", err)
		}
		return
	}

	val, _ := kf.GetValue(section, key)
	if val == "" {
		return
	}

	kf.DeleteKey(section, key)
	err = kf.SaveToFile(filename)
	if err != nil {
		logger.Warning("failed to save:", err)
	}
}
