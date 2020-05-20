package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"os"
)

func main() {
	app := iris.New()
	// 1. 通过程勋代码对应用进行全局配置
	app.Configure(iris.WithConfiguration(iris.Configuration{
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "mon,02,jan 2006 15:04:05 GMT",
		Charset:                           "utf-8",
	}))

	// 2. 通过读取tml配置文件读取服务配置
	// 注意：要在run方法运行之前执行
	app.Configure(iris.WithConfiguration(iris.TOML("/Users/dcw/Documents/go_project/src/irisDemo/5路由组以及iris配置/iris.tml ")))

	// 3. 通过读取yaml配置文件读取服务配置
	// 注意：同样要在run方法运行之前执行
	app.Configure(iris.WithConfiguration(iris.YAML("/Users/dcw/Documents/go_project/src/irisDemo/5路由组以及iris配置/iris.yml")))

	// 4. 通过jsos配置文件进行应用配置
	file, _ := os.Open("/Users/dcw/Documents/go_project/src/irisDemo/iris/5路由组以及iris配置/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Coniguration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(conf.Port)

	// 运行服务，端口监听
	app.Run(iris.Addr(":8009"))
}

type Coniguration struct {
	AppName string `json:"app_name"`
	Port    int    `json:"port"`
}
