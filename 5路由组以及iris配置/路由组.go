package main

import "github.com/kataras/iris"
import "github.com/kataras/iris/context"

func main() {
	// 路由组请求
	app := iris.New()

	userParty := app.Party("/users", func(context context.Context) {
		// 处理下一级请求
		context.Next()
	})

	// 路由组下面的下一级请求
	// ../users/register

	userParty.Get("/register", func(context context.Context) {
		app.Logger().Info("用户注册功能")
		context.HTML("</h1>用户注册功能</h1>")
	})

	// 路由组下面的下一级请求
	// ../users/login

	userParty.Get("/login", func(context context.Context) {
		app.Logger().Info("用户登录功能")
		context.HTML("</h1>用户登录功能</h1>")
	})

	// 不使用匿名函数,引用下面自己定义的函数 ,一个新的路由组：userRouter
	userRouter := app.Party("/admin", userMiddleWare)

	// 使用done方法
	userRouter.Done(func(context context.Context) {
		context.Application().Logger().Infof("response sent to" + context.Path())
	})

	userRouter.Get("/info", func(context context.Context) {
		context.HTML("<h1>用户信息</h1>")
		context.Next() // 手动显示调用done方法
	})
	userRouter.Get("/query", func(context context.Context) {
		context.HTML("<h1>用户信息</h1>")
	})

	app.Run(iris.Addr(":8003"), iris.WithoutServerError(iris.ErrServerClosed))

}

// 单独声明一个函数，不使用匿名函数
func userMiddleWare(context iris.Context) {
	context.Next()
}
