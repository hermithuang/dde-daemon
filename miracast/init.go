/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
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

package miracast

import (
	"sync"

	"pkg.deepin.io/dde/daemon/loader"
	"pkg.deepin.io/lib/log"
)

func init() {
	loader.Register(NewDaemon(logger))
}

type Daemon struct {
	*loader.ModuleBase
	wg sync.WaitGroup
}

func NewDaemon(logger *log.Logger) *Daemon {
	daemon := new(Daemon)
	daemon.ModuleBase = loader.NewModuleBase("miracast", daemon, logger)
	daemon.wg.Add(1)
	return daemon
}

func (d *Daemon) WaitEnable() {
	d.wg.Wait()
}

func (*Daemon) GetDependencies() []string {
	return []string{"network"}
}

var (
	_m     *Miracast
	logger = log.NewLogger("daemon/miracast")
)

func (d *Daemon) Start() error {
	defer d.wg.Done()
	if _m != nil {
		return nil
	}
	service := loader.GetService()

	m, err := newMiracast(service)
	if err != nil {
		logger.Error("Failed to new manager:", err)
		return err
	}
	_m = m

	err = service.Export(dbusPath, m)
	if err != nil {
		logger.Error("Failed to export:", err)
		_m.destroy()
		_m = nil
		return err
	}

	err = service.RequestName(dbusServiceName)
	if err != nil {
		logger.Error("Failed to request name:", err)
		_m.destroy()
		_ = service.StopExport(m)
		_m = nil
		return err
	}

	return nil
}

func (*Daemon) Stop() error {
	if _m == nil {
		return nil
	}
	_m.destroy()
	service := loader.GetService()
	_ = service.StopExport(_m)
	_m = nil
	return nil
}
