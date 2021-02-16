package controller

import (
	"going"
	"going/base/infra"
)

func init() {
	going.Prepare(func(initiator going.Initiator) {
		/*
			Common binding, default controller to path '/'.
			initiator.BindController("/",  &DefaultController{})
		*/

		// Middleware binding. Valid only for this controller.
		// If you need global middleware, pleas add in main.
		initiator.BindController("/", &Default{}, func(context going.Context) {
			worker := going.ToWorker(context)
			worker.Logger().Info("Hello middleware begin")
			context.Next()
			worker.Logger().Info("Hello middleware end")
		})
	})
}

type Default struct {
	Worker going.Worker
}

// Get handles the GET: / route.
func (c *Default) Get() going.Result {
	c.Worker.Logger().Info("I'm Controller")

	return &infra.JSONResponse{
		Object: "Hello",
	}
}
