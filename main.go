package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/mysql"
	_ "memplus_service/routers"
	"net/url"
	"os"
)
var globalSessions *session.Manager
const HOST = ""

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		orm.Debug = true
	}

	cnf, err := config.NewConfig("ini", "./conf/app.conf")

	if err != nil{
		os.Exit(-1)
	}
	mysqlConfig := cnf.String("mysql") + "?parseTime=true&loc=" +url.QueryEscape("Asia/Shanghai") //2018-12-27T14:55:20+08:00
	beego.BConfig.WebConfig.Session.SessionProvider = "mysql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = mysqlConfig
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600 * 24 // beego的session是定时删除的
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.DefaultRowsLimit = 100000000
    orm.RegisterDataBase("default", "mysql", mysqlConfig)
	beego.Run()
}
