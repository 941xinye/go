package main

import (
	"RESTful/methods"
	_ "RESTful/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hprose/hprose-golang/rpc"
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
	artivity := &methods.Artivity{}
	service := rpc.NewHTTPService()
	service.AddAllMethods(artivity)
	beego.Handler("/artivity", service)

	article := &methods.Article{}
	service1 := rpc.NewHTTPService()
	service1.AddAllMethods(article)
	beego.Handler("/article", service1)
	beego.Run()
}
