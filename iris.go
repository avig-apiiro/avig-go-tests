package main

import (
	"github.com/kataras/iris/v12"
)

func irismain() {
	app := iris.Default()

	app.Get("/user/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s", name)
	})
	v2 := app.Party("/v2")
	{
		v2.Post("/login", loginEndpoint)
		v2.Put("/submit", submitEndpoint)
		v2.Head("/read3", readEndpoint)
	}
	app.Listen(":8081")
}

func loginEndpoint(ctx iris.Context) {
	ctx.StatusCode(iris.StatusCreated)
}
func submitEndpoint(ctx iris.Context) {
	ctx.StatusCode(iris.StatusCreated)
}
func readEndpoint(ctx iris.Context) {
	ctx.StatusCode(iris.StatusCreated)
}
