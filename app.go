package going

import "going/internal"

var (
	// app is a singleton of internal.Application. It is never nil.
	app *internal.Application
)

func initApp() {
	app = internal.NewApplication()
}
