package main

import (
	_ "RESTful/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// orm.RegisterDataBase("default", "mysql", "root:@/seven_day?charset=utf8")
	// orm.RegisterDataBase("default", "mysql", "restful:Uft8ibTYTXE2u7oT@115.28.241.202/seven_day?charset=utf8")
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String(beego.AppConfig.String("runmode")+"::mysqluser")+":"+beego.AppConfig.String(beego.AppConfig.String("runmode")+"::mysqlpass")+"@"+beego.AppConfig.String(beego.AppConfig.String("runmode")+"::mysqlurls")+"/"+beego.AppConfig.String(beego.AppConfig.String("runmode")+"::mysqldb")+"?charset=utf8")
}

func main() {
	if beego.BConfig.RunMode == "dev" || beego.BConfig.RunMode == "localhost" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
