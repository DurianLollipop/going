package internal

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/host"
	"github.com/kataras/iris/v12/mvc"
)

type (
	// Dependency is a function which represents a dependency of the application.
	// A valid definition of dependency should like:
	//  func foo(ctx iris.Context) Bar {
	//    return Bar{}
	//  }
	//
	// Example:
	// If you want to inject a `Foo` into somewhere, you could written:
	//  func (ctx iris.Context) Foo {
	//    return NewFoo() // you'd replace NewFoo() with your own logic for creating Foo.
	//  }
	Dependency = interface{}
)

// Application represents a manager.
type Application struct {
	// prefixPath is the prefix of the application-level router.
	prefixPath string

	// controllerDependencies contains the dependencies that provided for each.
	// isis controller.
	controllerDependencies []interface{}

	// TODO
	// Here is bad smell I felt, so we shouldn't exported this member. Considering sets this member unexported in the future.
	// IrisApp is an iris application.
	IrisAPP *IrisApplication
}

// Logger
func (app *Application) Logger() *golog.Logger {
	return app.Iris().Logger()
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

func (app Application) BindController(relativePath string, controller IrisController, handlers ...IrisHandle) {
	router := app.SubRouter(relativePath, handlers...)

	mvcApp := mvc.New(router)
	mvcApp.Register(app.resolveDependencies()...)
	mvcApp.Handle(controller)
}

// SubRouter accepts a string with the route path, and one or more IrisHandler.
// SubRouter makes a new path by concatenating the application-level router's
// prefix and the specified route path, creates an IrisRouter with the new path
// and the those specified IrisHandler, and returns the IrisRouter.
func (app *Application) SubRouter(relativePath string, handlers ...IrisHandle) IrisParty {
	return app.Iris().Party(app.withPrefix(relativePath), handlers...)
}

// withPrefix returns a string with a path  prefixed by prefixPath.
func (app *Application) withPrefix(path string) string {
	return app.prefixPath + path
}

func (app *Application) resolveDependencies() []Dependency {
	var result []Dependency
	result = append(result, func(ctx IrisContext) (rt Worker) {
		rt = ctx.Values().Get(WorkerKey).(Worker)
		return
	})
	result = append(result, app.controllerDependencies...)
	return result
}
