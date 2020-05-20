package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
)

/*
Session 的使用和控制
*/

var USERNAME = "userName"
var ISLOGIN = "isLogin"

func main() {
	app := iris.New()
	sessionID := "mySession"
	//1.创建session并进行使用
	sess := sessions.New(sessions.Config{
		Cookie: sessionID,
	})

	// 用户登录功能
	app.Post("/login", func(context context.Context) {
		path := context.Path()
		app.Logger().Info("请求Path", path)
		userName := context.PostValue("name")
		passwd := context.PostValue("pwd")
		if userName == "davie" && passwd == "pwd123" {
			session := sess.Start(context)
			session.Set(USERNAME, userName)
			session.Set(ISLOGIN, true)
			context.WriteString("账户登录成功")
		} else {
			session := sess.Start(context)
			session.Set(ISLOGIN, false)
			context.WriteString("账户登录失败，请重新尝试")
		}

	})

	// 用户推出登录功能
	app.Get("/logout", func(context context.Context) {
		path := context.Path()
		app.Logger().Info("退出登录path:", path)
		session := sess.Start(context)
		//删除session
		session.Delete(ISLOGIN)
		context.WriteString("退出登录成功")

	})


	// session 查询
	app.Get("/query", func(context context.Context) {
		path := context.Path()
		app.Logger().Info("查询信息 path:",path)
		session := sess.Start(context)
		isLogin ,err := session.GetBoolean(ISLOGIN)
		if err !=nil{
			context.WriteString("账户未登录，请先登录")
			return
		}
		if isLogin{
			app.Logger().Info("账户已登录")
			context.WriteString("账户已登录")
		} else {
			app.Logger().Info("账户未登录")
			context.WriteString("账户未登录")
		}

	})










}
