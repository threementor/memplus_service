package main

import (
	_ "memplus_service/routers"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/config"
	"os"
)
var globalSessions *session.Manager
const HOST = ""

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	cnf, err := config.NewConfig("ini", "./conf/app.conf")

	if err != nil{
		os.Exit(-1)
	}

	mysql_config := cnf.String("mysql")
	beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = mysql_config
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.DefaultRowsLimit = 100000000
    orm.RegisterDataBase("default", "mysql", mysql_config)
	orm.Debug = cnf.DefaultBool("debug", false)

	beego.Run()
}
