package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	// 返回string
	// 返回html
	//返回 json
	app := iris.New()
	app.Run(iris.Addr(":8082"))
	app.Get("/getJson", func(context context.Context) {
		context.JSON(iris.Map{"message": "helloworld", "requestcode": 200})
	})
}
