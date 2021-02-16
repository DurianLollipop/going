package going

import (
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
	Run(serve iris.Runner, c iris.Configuration)
}

// NewApplication
func NewApplication() Application {
	return app
}
