package internal

import (
	"github.com/kataras/iris/v12/context"
	"sync"
)

var (
	globalApp     *Application
	globalAppOnce sync.Once
	prepares      []func(Initiator)
)

// Initiator.
type Initiator interface {
	BindController(relativePath string, controller interface{}, handlers ...context.Handler)
}

// Prepare app. BindController or app.BindControllerByParty.
func Prepare(f func(Initiator)) {
	prepares = append(prepares, f)
}
