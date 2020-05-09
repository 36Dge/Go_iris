package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {

	app := iris.New()

	//1. handle方式处理请求
	app.Handle("GET", "/userinfo", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
	})
	//2.Handle方式处理请求post
	app.Handle("POST", "/postcommit", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		context.HTML(path)
	})

	//3. 请求路径中设置变量 正则表达式 Get方式
	app.Get("/weather/{data}/{city}", func(context context.Context) {
		path := context.Path()
		// 如何获取get请求中的正则表达式方法
		data := context.Params().Get("data")
		city := context.Params().Get("city")

		context.WriteString(path + data + city)

	})

	app.Get("/hello/{name}", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)

		//或取正则表达式内容值
		name := context.Params().Get("name")
		context.HTML(name)

	})

	// 自定义正则表达式路由请求

	app.Get("/api/users/{userid:uint64}", func(context context.Context) {
		userID, err := context.Params().GetUint("useri")
		if err != nil {
			//设置请求状态码,状态码可以自定义
			if err != nil {
				//设置请求状态码，状态码可以自定义
				context.JSON(map[string]interface{}{
					"requestcode": 201,
					"message":     "bad ruquest",
				})
				return
			}

			context.JSON(map[string]interface{}{
				"requestcode": 200,
				"user_id":     userID,
			})
		}
	})

	app.Get("/api/users/{isLogin:bool}", func(context context.Context) {
		isLogin, err := context.Params().GetBool("isLogin")
		if err != nil {
			context.StatusCode(iris.StatusNonAuthoritativeInfo)
			return
		}
		if isLogin {
			context.WriteString("已经登")
		} else {
			context.WriteString("未登录")
		}
	})

	//
	app.Run(iris.Addr("：8083"))

}
