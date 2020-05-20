package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	// 路由组请求
	app := iris.New()

	// 用户模块
	//路由组请求
	userParty := app.Party("/users", func(context context.Context) {
		//处理下一级请求
		context.Next()
	})

	// 路由组下一级请求
	userParty.Get("/register", func(context context.Context) {
		app.Logger().Info("用户注册功能")
		context.HTML("<h1>用户注册功能</h1>")
	})

	userParty.Get("/login", func(context context.Context) {
		app.Logger().Info("用户登录功能")
		context.HTML("<h1>用户登录功能</h1>")
	})

	// 设置路由组中间件，将第二个匿名函数自己定义一个函数
	userRouter := app.Party("/admin", userMiddleWare)

	// dong 方法
	userRouter.Done(func(context context.Context) {
		context.Application().Logger().Infof("respose sent to " + context.Path())
	})

	userRouter.Get("/info", func(context context.Context) {
		context.HTML("用户信息")
		context.Next()
	})

	userRouter.Get("/query", func(context context.Context) {
		context.HTML("查询信息")
	})

	app.Run(iris.Addr(":8003"), iris.WithoutServerError(iris.ErrServerClosed))

}

func userMiddleWare(context iris.Context) {
	context.Next()
}
