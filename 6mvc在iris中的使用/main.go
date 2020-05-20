package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()

	//设置自定义控制器
	mvc.New(app).Handle(new(UserController))
	//路由组的mvc处理
	mvc.Configure(app.Party("/user"), func(context *mvc.Application) {
		context.Handle(new(UserController))
	})

}

// 自动匹配对应控制器的方法
func (uc *UserController) Get() string {
	iris.New().Logger().Info("Get请求")
	return "hello world"
}

func (uc *UserController) Post() {
	iris.New().Logger().Info("post请求")
}

// 匹配url
func (uc *UserController) GetInfo() mvc.Result {
	iris.New().Logger().Info("get请求，请求路径为info")
	return mvc.Response{
		Object: map[string]interface{}{
			"code":    1,
			"message": "请求成功",
		},
	}

}

// 自定义控制器
type UserController struct {
	// to do
}

func (uc *UserController) BeforeActivation(a mvc.BeforeActivation) {
	a.Handle("GET", "/query", "UserInfo")
}

func (uc *UserController) UserInfo() mvc.Result {
	//to do
	iris.New().Logger().Info("user info query")
	return mvc.Response{}
}
