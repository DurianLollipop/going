package internal

import "github.com/kataras/iris/v12"

// Application represents a manager.
type Application struct {
	// TODO
	// Here is bad smell I felt, so we shouldn't exported this member. Considering sets this member unexported in the future.
	// IrisApp is an iris application.
	IrisAPP *IrisApplication
}

// NewApplication creates an instance of Application exactly once while the program is running.
func NewApplication() *Application {
	globalAppOnce.Do(func() {
		globalApp = new(Application)
		globalApp.IrisAPP = iris.New()
	})
	return globalApp
}
