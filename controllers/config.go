package controllers

import "github.com/astaxie/beego"


func GetHostName() string {
	return beego.AppConfig.String("host")
}