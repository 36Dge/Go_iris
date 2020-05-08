package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {

	// 创建application对象
	app := iris.New()

	app.Get("/getRequest", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
	})
	//1.处理get请求
	app.Get("/userpath", func(context context.Context) {
		//获取path
		path := context.Path()
		app.Logger().Info(path)
		//写入返回数据，string类型
		context.WriteString("请求数据" + path)

	})

	//2.处理get请求，并接收参数
	app.Get("/userinfo", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		//获取get请求的所携带的参数
		userName := context.URLParam("username")
		app.Logger().Info(userName)

		pwd := context.URLParam("pwd")
		app.Logger().Info(pwd)

		// 返回数据，返回HTML格式数据
		context.HTML("<h1>" + userName + "," + "</h1>")

	})

	// 3 .处理post请求以及所携带的参数

	app.Post("/postLogin", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		//获取post请求携带的数据
		name := context.PostValue("name")
		pwd := context.PostValue("pwd")
		app.Logger().Info(name)
		app.Logger().Info(pwd)
		context.HTML(name)

	})

	// 4.处理post请求以及json类型的数据参数
	app.Post("/postJson", func(context context.Context) {
		//1.path
		path := context.Path()
		app.Logger().Info("请求URL", path)

		//2.postjson数据解析
		var person Person
		//context.ReadJSON()
		if err := context.ReadJSON(&person); err != nil {
			panic(err.Error())
		}
		context.Writef("received",person)

	})

	//端口监听
	app.Run(iris.Addr(":8081"))

}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
