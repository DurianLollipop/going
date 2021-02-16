package internal

const (
	// WorkerKey specify where is the Worker should be located in Context
	WorkerKey = "STORE-WORKER-KEY"
)

// Worker describes a global context which use to share the internal component
// (i.e infrastructure, transaction, logger and so on) with middleware,
// controller, domain service and etc.
type Worker interface {

	// Logger returns current Logger.
	Logger() Logger
}
