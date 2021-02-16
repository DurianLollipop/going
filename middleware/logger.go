package middleware

import (
	"github.com/kataras/golog"
	"going"
)

// Logger.
type Logger struct {
	traceId   string
	traceName string
}

// Info will print when Logger's Level is info or debug.
func (l *Logger) Info(v ...interface{}) {
	trace := l.traceFiled()
	v = append(v, trace)
	going.Logger().Info(v...)
}

// traceFiled
func (l Logger) traceFiled() golog.Fields {
	return golog.Fields{l.traceName: l.traceId}
}
