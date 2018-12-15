package controllers
import (
	"github.com/astaxie/beego"
	"memplus_service/models"
	"errors"
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
