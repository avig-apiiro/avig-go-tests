package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	app.Get("/V3/user/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s", name)
	})
	v2 := app.Party("/v2")
	{
		v2.Post("/V3/login", loginEndpoint)
	}
	app.Listen(":8081")
}

func loginEndpoint(ctx iris.Context) {
	ctx.StatusCode(iris.StatusCreated)
}
