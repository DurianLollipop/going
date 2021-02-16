package main

import (
	"github.com/kataras/iris/v12"
	"going"
)

func main() {
	application := going.NewApplication()
	addrRunner := application.NewRunner(":8080")
	application.Run(addrRunner, iris.Configuration{})
}
