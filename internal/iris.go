package internal

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

type (
	// IrisApplication is type alias to iris.Application.
	IrisApplication = iris.Application

	// IrisContext is a type alias iris.Context
	IrisContext = iris.Context

	// IrisRunner is a type alias to iris.Runner.
	IrisRunner = iris.Runner

	// IrisConfiguration is a type alias to iris.Configuration.
	IrisConfiguration = iris.Configuration

	// IrisParty is a type alias to iris,Party.
	IrisParty = iris.Party

	// IrisHandle is a type alias to context.Handle
	IrisHandle = context.Handler

	// IrisController represents a standard iris controller.
	IrisController = interface{}
)
