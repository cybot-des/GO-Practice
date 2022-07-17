package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.Get("/", callBack)
	app.Listen(":9000")
}

func callBack(ctx iris.Context) {
	ctx.JSON("CallBack Function Called")
}
