package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func main() {
	// 1. 创建数据库引擎对象
	engine, err := xorm.NewEngine("mysql", "root:e13212ss@/mysql?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	//2.数据库引擎关闭
	defer engine.Close()

	//数据库引擎设置
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)

	//将person结构体同步到数据库中并创建person表
	engine.Sync2(new(Person))

}

type Person struct {
	Ager int
	Name string
}
