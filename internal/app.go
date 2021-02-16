package internal

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/host"
)

// Application represents a manager.
type Application struct {
	// TODO
	// Here is bad smell I felt, so we shouldn't exported this member. Considering sets this member unexported in the future.
	// IrisApp is an iris application.
	IrisAPP *IrisApplication
}

func (app *Application) Run(serve IrisRunner, conf IrisConfiguration) {
	app.Iris().Run(serve, iris.WithConfiguration(conf))
}

func (app *Application) NewRunner(addr string, configurators ...host.Configurator) iris.Runner {
	return iris.Addr(addr, configurators...)
}

// NewApplication creates an instance of Application exactly once while the program is running.
func NewApplication() *Application {
	globalAppOnce.Do(func() {
		globalApp = new(Application)
		globalApp.IrisAPP = iris.New()
	})
	return globalApp
}

// Iris.
func (app *Application) Iris() *IrisApplication {
	return app.IrisAPP
}
