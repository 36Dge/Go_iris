package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func main() {

	//1.创建数据库引擎
	engine, err := xorm.NewEngine("mysql", "root:e13212ss@/Test?charset=utf8")

	if err != nil {
		panic(err.Error())
	}
	//设置名称映射规则;
	engine.SetMapper(core.SnakeMapper{})
	engine.Sync2(new(UserTabel))

	engine.SetMapper(core.SameMapper{})
	engine.Sync2(new(StudentTabel))

	engine.SetMapper(core.GonicMapper{})
	engine.Sync2(new(PersonTable))

	personEmpty, err := engine.IsTableEmpty(new(PersonTable))
	if err != nil {
		panic(err.Error())
	}
	if personEmpty {
		fmt.Println("人员表是空的")
	} else {
		fmt.Println("人员表不为空")
	}

}

type UserTabel struct {
	UserId   int64  `xorm:"pk autoincr"`
	UserName string `xorm:"varchar(32)"`
	UserAge  int64  `xorm:"default 1"`
	UserSex  int64  `xorm:default 0`
}

type StudentTabel struct {
	Id          int64  `xorm:"pk autoincr"`
	StudentName string `xorm:"varchar(24)"`
	StudenAge   int    `xorm:"int default 0"`
	StudentSex  int    `xorm:"index"`
}

//人员结构体表
type PersonTable struct {
	ID         int64     `xorm:"pk autoincr"`
	PersonName string    `xorm:"varchar(24)"`
	PersonAge  int       `xorm:int "default 0"`
	PersonSex  int       `xorm:"notnull"`
	City       CityTable `xorm:-`
}

type CityTable struct {
	CityName      string
	CityLongitude float64
	Citylatitude  float64
}
