package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
)

var (
	USERNAME = "userName"
	ISLOGIN  = "isLogin"
)

/*
Session 的使用和控制
*/

func main() {

	app := iris.New()
	sessionID := "mySession"
	// 1. 创建Session并使用
	sess := sessions.New(sessions.Config{
		Cookie: sessionID,
	})

	// 2.Session的存储和使用
	// 用户登录功能

	app.Post("/login", func(context context.Context) {
		path := context.Path()
		app.Logger().Info("请求Path:", path)
		userName := context.PostValue("name")
		passWd := context.PostValue("pwd")
		if userName == "davie" && passWd == "pwd123" {
			session := sess.Start(context)
			// 用户名
			session.Set(USERNAME, userName)
			// 登录状态
			session.Set(ISLOGIN, true)
			context.WriteString("账户登录成功")

		} else {
			session := sess.Start(context)
			session.Set(ISLOGIN, false)
			context.WriteString("账户登录失败，请重新尝试")
		}
	})

	// 用户推出登录功能

	app.Get("./logout", func(context context.Context) {
		path := context.Path()
		app.Logger().Info("退出登录Path:", path)
		session := sess.Start(context)
		session.Delete(ISLOGIN)
		session.Delete(USERNAME)
		context.WriteString("退出登录成功")

	})

	//用户查询

	app.Get("./query", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		session := sess.Start(context)
		// GetBoolean()
		isLogin, err := session.GetBoolean(ISLOGIN)
		if err != nil {
			context.WriteString("账户未登录，请先登录")
			return
		}

		if isLogin {
			app.Logger().Info("账户已登录")
			context.WriteString("账户已经登录")
		} else {
			app.Logger().Info("账户未登录")
			context.WriteString("账户未登录")
		}

	})

	// 3. Session与数据库结合使用
/*
	db, err := boltdb.New(".sessions.db", 0600)
	if err != nil {
		panic(err.Error())
	}
	// 程序中断时，将数据库关闭
	iris.RegisterOnInterrupt(func() {
		defer db.Close()
	})

	// session和db绑定
	sess.UseDatabase(db)

	app.Run(iris.Addr(":8009"))

 */
}

