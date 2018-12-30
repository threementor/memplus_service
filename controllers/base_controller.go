package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"memplus_service/models"
)

type LoginRequirePrepare interface {
	LoginRequirePrepare()
}


type BaseController struct {
	beego.Controller
	LoginRequire bool
}


func (this *BaseController) Prepare() {
	uid := this.GetSession("uid")

	this.Data["isLogin"] = false

	if uid != nil{
		println("uid is not null", uid)
		uid_int, ok := uid.(int)
		if ok{
			user, err := models.GetUsersById(uid_int)
			this.Data["isLogin"] = err == nil
			this.Data["user"] = user
		}
	}else{
		println("uid is null")
	}
	if app, ok := this.AppController.(LoginRequirePrepare); ok {
		app.LoginRequirePrepare()
	}
}

func (this *BaseController) GetUser() (*models.User, error){
	user, ok := this.Data["user"]
	if !ok{
		return nil, errors.New("user not in data")
	}
	if u, ok := user.(*models.User); ok{
		return u, nil
	}else{
		return nil, errors.New("cover user fail")
	}
}


func (this *BaseController) SendSuccess(data interface{}){
	this.Data["json"] = map[string]interface{}{"code": 200, "data": data}
	this.ServeJSON()
}


func (this *BaseController) SendError(err error){
	this.Data["json"] = map[string]interface{}{"code": -1, "msg": err.Error()}
	this.ServeJSON()
}


type LoginReqireController struct{
	BaseController
}

func (this *LoginReqireController) LoginRequirePrepare(){
	isLogin, ok := this.Data["isLogin"]
	if ok{
		isLoginBool := isLogin.(bool)
		if !isLoginBool{
			this.Abort("401")
		}
	}
}

