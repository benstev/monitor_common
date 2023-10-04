package app

import (
	"context"

	"github.com/benstev/monitor_common/service"

	"github.com/benstev/monitor_common/initializers"
)

type AppBuilder = func(context.Context) (*Application, error)

type Dependencies interface {
	StartServices()
	StopServices()
}

type Application struct {
	DaprService *service.Service
	Container   Dependencies
}

func (a *Application) Start(ctx context.Context, cli bool) {
	if cli {
		return
	}

	a.DaprService.Start()
	a.Container.StartServices()
}

func (a *Application) Stop() (err error) {
	if err := a.DaprService.Stop(); err != nil {
		return err
	}
	a.Container.StopServices()
	return nil
}

func InitializeApplication(ctx context.Context, builder AppBuilder) (*Application, error) {
	initializers.InitializeEnvs()
	initializers.InitializeLogs()
	return builder(ctx)
}
