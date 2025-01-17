package image_effect

import (
	"sync"

	"pkg.deepin.io/dde/daemon/loader"
	"pkg.deepin.io/lib/log"
)

func init() {
	loader.Register(newModule())
}

type Module struct {
	ie *ImageEffect
	*loader.ModuleBase
	wg sync.WaitGroup
}

func (m *Module) WaitEnable() {
	m.wg.Wait()
}

func (m *Module) GetDependencies() []string {
	return nil
}

func (m *Module) Start() error {
	defer m.wg.Done()
	if m.ie != nil {
		return nil
	}

	var err error
	m.ie, err = start()
	if err != nil {
		return err
	}

	return nil
}

func (m *Module) Stop() error {
	// TODO
	return nil
}

const moduleName = "image_effect"

var logger = log.NewLogger("daemon/" + moduleName)

func newModule() *Module {
	m := &Module{}
	m.ModuleBase = loader.NewModuleBase(moduleName, m, logger)
	m.wg.Add(1)
	return m
}

func start() (*ImageEffect, error) {
	logger.Debug("module image_effect start")
	ie := newImageEffect()
	service := loader.GetService()
	ie.service = service
	err := service.Export(dbusPath, ie)
	if err != nil {
		return nil, err
	}

	err = service.RequestName(dbusServiceName)
	if err != nil {
		return nil, err
	}

	return ie, nil
}
