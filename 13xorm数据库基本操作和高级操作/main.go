package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
)

func main() {

	//1.创建数据库引擎
	engine, err := xorm.NewEngine("mysql", "root:e13212ss@/Test?charset=utf8")

	if err != nil {
		panic(err.Error())
	}

	// 2.条件查询
	// 1. Id查询 Get方法，单条记录
	var person PersonTable
	engine.ID(1).Get(&person)
	fmt.Println(person.PersonName)
	fmt.Println()
	//2.where多条件查询  使用find
	var person1 []PersonTable
	engine.Where("person_age = ? and person_sex = ?", 26, 2).Find(&person1)

	//3.And条件查询
	var persons []PersonTable
	err = engine.Where("person_age = ?", 26).And("person_sex = ?", 2).Find(&persons)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(persons)
	fmt.Println()

	//4. or 条件查询
	var perosnArr []PersonTable
	err = engine.Where("person_age = ?", 26).Or("person_sex = ?", 1).Find(&perosnArr)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(perosnArr)
	fmt.Println()

	//5.原生SQL语句查询支持 like语法
	var personNative []PersonTable
	err = engine.SQL("select * from perosn_table where person_name like '%t'").Find(&personNative)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(personNative)
	fmt.Println()

	//6.排序条件查询
	var personOrderBy []PersonTable
	engine.OrderBy("person_age desc").Find(&personOrderBy)
	fmt.Println(personOrderBy)
	fmt.Println()

	//7.查询特定字段
	var personCols []PersonTable
	engine.Cols("person_name", "person_age").Find(&personCols)
	for _, col := range personCols {
		fmt.Println(col)
	}
	//三、增加记录操作
	personInsert := PersonTable{
		PersonName: "hello",
		PersonAge:  18,
		PersonSex:  1,
	}
	rowNum, err := engine.Insert(&personInsert)
	fmt.Println(rowNum)
	fmt.Println()

	//四、删除操作

	rowNum, err = engine.Delete(&personInsert)
	fmt.Println(rowNum)
	fmt.Println()

	//五、更新操作
	rowNum, err = engine.ID(7).Update(&personInsert)
	fmt.Println(rowNum)
	fmt.Println()

	//六、统计功能count
	count, err := engine.Count(new(PersonTable))
	fmt.Println("persontable表总记录条数", count)

	//七、事物操作
	//...

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
