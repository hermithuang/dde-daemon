/*
 * Copyright (C) 2019 ~ 2020 Uniontech Software Technology Co.,Ltd.
 *
 * Author:     zhihsian <i@zhihsian.me>
 *
 * Maintainer: zhihsian <i@zhihsian.me>
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

package bluetooth

import (
	obex "github.com/linuxdeepin/go-dbus-factory/org.bluez.obex"
	dbus "pkg.deepin.io/lib/dbus1"
	"pkg.deepin.io/lib/dbusutil"
)

const (
	obexAgentDBusPath      = dbusPath + "/ObexAgent"
	obexAgentDBusInterface = "org.bluez.obex.Agent1"
)

type obexAgent struct {
	b *Bluetooth

	service     *dbusutil.Service
	sigLoop     *dbusutil.SignalLoop
	obexManager *obex.Manager

	methods *struct {
		AuthorizePush func() `in:"transferPath" out:"filename"`
		Cancel        func()
	}
}

func (*obexAgent) GetInterfaceName() string {
	return obexAgentDBusInterface
}

func newObexAgent(service *dbusutil.Service, bluetooth *Bluetooth) *obexAgent {
	return &obexAgent{
		b:       bluetooth,
		service: service,
	}
}

func (a *obexAgent) init() {
	sessionBus := a.service.Conn()
	a.obexManager = obex.NewManager(sessionBus)
	a.registerAgent()
}

func (a *obexAgent) registerAgent() {
	err := a.obexManager.RegisterAgent(0, obexAgentDBusPath)
	if err != nil {
		logger.Error("failed to register obex agent:", err)
	}
}

func (a *obexAgent) unregisterAgent() {
	err := a.obexManager.UnregisterAgent(0, obexAgentDBusPath)
	if err != nil {
		logger.Error("failed to unregister obex agent:", err)
	}
}

// AuthorizePush 用于请求用户接收文件
func (a *obexAgent) AuthorizePush(transferPath dbus.ObjectPath) (string, *dbus.Error) {
	transfer, err := obex.NewTransfer(a.service.Conn(), transferPath)
	if err != nil {
		logger.Error("failed to new transfer:", err)
		return "", dbusutil.ToError(err)
	}

	filename, err := transfer.Name().Get(0)
	if err != nil {
		logger.Warning("failed to get filename:", err)
		return "", dbusutil.ToError(err)
	}

	return filename, nil
}

// Cancel 用于在客户端取消发送文件时取消文件传输请求
func (a *obexAgent) Cancel() *dbus.Error {
	return nil
}
