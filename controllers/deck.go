package controllers

import (
	"encoding/json"
	"errors"
	"memplus/models"
	"strconv"
	"strings"
	"fmt"
)

// KlgDirController operations for KlgDir
type KlgDirController struct {
	LoginReqireController
}

// URLMapping ...
func (c *KlgDirController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create KlgDir
// @Param	body		body 	models.KlgDir	true		"body for KlgDir content"
// @Success 201 {int} models.KlgDir
// @Failure 403 body is empty
// @router / [post]
func (c *KlgDirController) Post() {
	var v models.KlgDir
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		user, err := c.GetUser()
		if err != nil{
			c.Data["json"] = err.Error()
		}else{
			v.UserId = user.Id
			if _, err := models.AddKlgDir(&v); err == nil {
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = v
			} else {
				c.Data["json"] = err.Error()
			}
		}


	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}


// GetTasks ...
// @Title Get GetTasks
// @Description get GetTasks  by KlgDir id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Task
// @router /:id/ready_tasks [get]
func (c *KlgDirController) GetReadyTasks(){
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	dir, err := models.GetKlgDirById(id)
	if err != nil{
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	user, _ := c.GetUser()
	tasks, err := models.GetReadyTasks(dir, user)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = tasks
	}
	c.ServeJSON()
}


// GetOne ...
// @Title Get One
// @Description get KlgDir by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.KlgDir
// @Failure 403 :id is empty
// @router /:id [get]
func (c *KlgDirController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetKlgDirById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get KlgDir
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.KlgDir
// @Failure 403
// @router / [get]
func (c *KlgDirController) GetAll() {
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

	user, err := c.GetUser()
	if err == nil{
		query["user_id"] = fmt.Sprintf("%v", user.Id)
		l, err := models.GetAllKlgDir(query, fields, sortby, order, offset, limit)
		if err != nil {
			if err.Error() == "<QuerySeter> no row found"{
				c.Data["json"] = [][]string{}
			}else{
				c.Data["json"] = err.Error()
			}
		} else {
			c.Data["json"] = l
		}

	}else{
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}


// GetAll ...
// @Title Get Root Dirs
// @Description get KlgDir
// @Success 200 {object} models.KlgDir
// @Failure 403
// @router /roots [get]
func (c *KlgDirController) GetRootDirs() {

	user, err := c.GetUser()
	if err != nil{
		c.Data["json"] = err.Error()
		c.ServeJSON()
		c.Abort("401")
		return
	}
	dirs, err := models.GetRootDirs(user)

	if err != nil {
		if err.Error() == "<QuerySeter> no row found"{
			c.Data["json"] = [][]string{}
		}else{
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = dirs
	}
	c.ServeJSON()
}


// Put ...
// @Title Put
// @Description update the KlgDir
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.KlgDir	true		"body for KlgDir content"
// @Success 200 {object} models.KlgDir
// @Failure 403 :id is not int
// @router /:id [put]
func (c *KlgDirController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.KlgDir{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateKlgDirById(&v); err == nil {
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
// @Description delete the KlgDir
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *KlgDirController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteKlgDir(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}