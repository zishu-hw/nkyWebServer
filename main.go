package main

import (
	"nkyWebServer/models"
	_ "nkyWebServer/routers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"

	"github.com/astaxie/beego"
)

func init() {
	// 注册数据库，最大连接数30
	err := orm.RegisterDataBase("default", "mysql", "root:zy0802@tcp(localhost:3306)/db_test?charset=utf8", 30)
	if err != nil {
		beego.Error(err)
	}
	models.RegisterDB()
}

func main() {
	toolbox.StartTask()
	defer toolbox.StopTask()

	beego.Run()
}
