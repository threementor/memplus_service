package controllers

import (
	"encoding/json"
	"errors"
	"memplus_service/models"
	"strconv"
	"strings"
	"fmt"
)

// UserController operations for Users
type UserController struct {
	BaseController
}


func (c *UserController) InitFirst() {
	c.LoginRequire = false
}


// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Login", c.Login)
	c.Mapping("Logout", c.Logout)
}


// Post ...
// @Title Post
// @Description create Users
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 201 {int} models.Users
// @Failure 403 body is empty
// @router /reg [post]
func (c *UserController) Post() {
	var v models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if uid, err := models.AddUsers(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.SetSession("uid", int(uid))
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Users by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Users
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUsersById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Users
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Users
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUsers(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Users
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 200 {object} models.Users
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.User{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUsersById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Users
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUsers(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Param	username	query	string	true	"username"
// @Param	password	query	string	true	"password"
// @Success 200 {string}  success or fail!
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	fmt.Println(username, password)
	user, err := models.Login(username, password)
	if err == nil {
		u.SetSession("uid", user.Id)
		u.Data["json"] = map[string]interface{}{"success": true, "userName": username}
	} else {
		u.Data["json"] = map[string]interface{}{"success": false, "msg": err.Error()}
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.DestroySession()
	u.Data["json"] = map[string]interface{}{"success": true}
	u.ServeJSON()
}

// @Title status
// @Description user status
// @Success 200 
// @router /status/ [get]
func (u *UserController) Status() {
	user := u.Data["user"]
	if user == nil{
		u.Data["json"] = map[string]interface{}{"success": false}
	}else{
		u.Data["json"] = map[string]interface{}{"success": true, "userName": user.(*models.User).Name}
	}
	u.ServeJSON()
}

// @Title cahnge password
// @Description
// @Success 200
// @router /change_pwd/ [get]
func (u *UserController) ChangePwd() {
	newPwd := u.GetString("newPwd")
	user, _ := u.Data["user"].(*models.User)
	isLogin := u.Data["isLogin"].(bool)
	if len(newPwd) < 6{
		u.Data["json"] = map[string]interface{}{"success": false, "msg": "密码不能少于6位"}
	}else if isLogin{
		u.Data["json"] = map[string]interface{}{"success": false, "msg": "请登录"}
	}else{
		err := models.ChangePassword(user, newPwd)
		if err != nil{
			u.Data["json"] = map[string]interface{}{"success": false, "msg": err.Error()}
		}else{
			u.Data["json"] = map[string]interface{}{"success": true}
		}
	}
	u.ServeJSON()
}

