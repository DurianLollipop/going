package going

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"going/internal"
)

func init() {
	initApp()
	//initConfigurator()
}

type (
	// IrisResult represents an type to hero.Result.
	IrisResult = hero.Result

	// IrisContext represents an type alias to iris.Context
	IrisContext = iris.Context
)

type (
	// Worker.
	Worker = internal.Worker

	// Initiator
	Initiator = internal.Initiator

	// Result is the controller return type.
	Result = IrisResult

	// Context is the context type.
	Context = IrisContext
)

// Prepare
func Prepare(f func(Initiator)) {
	internal.Prepare(f)
}
