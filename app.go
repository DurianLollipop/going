package going

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/host"
	"going/internal"
)

var (
	// app is a singleton of internal.Application. It is never nil.
	app *internal.Application
)

func initApp() {
	app = internal.NewApplication()
}

// Application represents an application.
type Application interface {
	NewRunner(addr string, configurators ...host.Configurator) iris.Runner
	Iris() *iris.Application
	Logger() *golog.Logger
	Run(serve iris.Runner, c iris.Configuration)
}

// NewApplication
func NewApplication() Application {
	return app
}

func Logger() *golog.Logger {
	return app.Logger()
}

// ToWorker proxy a call to WorkerFromCtx.
func ToWorker(ctx Context) Worker {
	return WorkerFormCtx(ctx)
}

// WorkerFromCtx extracts Worker from a Context.
func WorkerFormCtx(ctx Context) Worker {
	if result, ok := ctx.Values().Get(internal.WorkerKey).(Worker); ok {
		return result
	}
	return nil
}
