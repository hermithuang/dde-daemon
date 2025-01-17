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

// Manage desktop appearance
package appearance

import (
	"errors"
	"sync"
	"time"

	"github.com/godbus/dbus"
	"pkg.deepin.io/dde/daemon/appearance/background"
	"pkg.deepin.io/dde/daemon/loader"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/log"
)

var (
	_m     *Manager
	logger = log.NewLogger("daemon/appearance")
)

type Module struct {
	*loader.ModuleBase
	wg sync.WaitGroup
}

func init() {
	background.SetLogger(logger)
	loader.Register(NewModule(logger))
}

func (m *Module) WaitEnable() {
	m.wg.Wait()
}

func NewModule(logger *log.Logger) *Module {
	var d = new(Module)
	d.ModuleBase = loader.NewModuleBase("appearance", d, logger)
	d.wg.Add(1)
	return d
}

func HandlePrepareForSleep(sleep bool) {
	if _m == nil {
		return
	}
	if sleep {
		return
	}
	cfg, err := doUnmarshalWallpaperSlideshow(_m.WallpaperSlideShow.Get())
	if err == nil {
		for monitorSpace := range cfg {
			if cfg[monitorSpace] == wsPolicyWakeup {
				_m.autoChangeBg(monitorSpace, time.Now())
			}
		}
	}
}

func (*Module) GetDependencies() []string {
	return []string{}
}

func (*Module) start() error {
	service := loader.GetService()

	_m = newManager(service)
	err := _m.init()
	if err != nil {
		logger.Warning(err)
		return err
	}

	err = service.Export(dbusPath, _m, _m.syncConfig)
	if err != nil {
		_m.destroy()
		return err
	}

	err = service.Export(backgroundDBusPath, _m.bgSyncConfig)
	if err != nil {
		return err
	}

	so := service.GetServerObject(_m)
	err = so.SetWriteCallback(_m, propQtActiveColor, func(write *dbusutil.PropertyWrite) *dbus.Error {
		value, ok := write.Value.(string)
		if !ok {
			return dbusutil.ToError(errors.New("type is not string"))
		}
		err = _m.setQtActiveColor(value)
		return dbusutil.ToError(err)
	})
	if err != nil {
		return err
	}

	err = service.RequestName(dbusServiceName)
	if err != nil {
		_m.destroy()
		err = service.StopExport(_m)
		if err != nil {
			return err
		}
		return err
	}

	err = _m.syncConfig.Register()
	if err != nil {
		logger.Warning("failed to register for deepin sync", err)
	}

	err = _m.bgSyncConfig.Register()
	if err != nil {
		logger.Warning("failed to register for deepin sync", err)
	}

	go _m.listenCursorChanged()
	go _m.handleThemeChanged()
	_m.listenGSettingChanged()
	return nil
}

func (m *Module) Start() error {
	defer m.wg.Done()
	if _m != nil {
		return nil
	}
	return m.start()
}

func (*Module) Stop() error {
	if _m == nil {
		return nil
	}

	_m.destroy()
	service := loader.GetService()
	err := service.StopExport(_m)
	if err != nil {
		return err
	}
	_m = nil
	return nil
}
