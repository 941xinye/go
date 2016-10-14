package main

import (
	_ "RESTful/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String(beego.AppConfig.String("runmode")+"::mysqluser")+":"+beego.AppConfig.String(beego.AppConfig.String("runmode")+"::mysqlpass")+"@"+beego.AppConfig.String(beego.AppConfig.String("runmode")+"::mysqlurls")+"/"+beego.AppConfig.String(beego.AppConfig.String("runmode")+"::mysqldb")+"?charset=utf8")
}

func main() {
	if beego.BConfig.RunMode == "dev" || beego.BConfig.RunMode == "localhost" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
