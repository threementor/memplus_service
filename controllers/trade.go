package controllers

import (
	"encoding/json"
	"errors"
	"memplus/models"
	"strconv"
	"strings"
	"github.com/smartwalle/alipay"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"

)


const appId = "2018101561694278"
const aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApNGh9/0sQqUAdBboAN5jdIXSSmQMBBC7GTKC6f8a4SUTW5I2cxtT6cEgZ63E4MBXmTfyLUYzoySVvY4Hq+4AA/EC4Fc6cxpHOrZtqltxiBUJcoclxUgURi/beR/q7rU+KAAN1uoOYsm8e7cfKR3LCz0v9773ZSTILTl9dkVRGcGySGUmgjQi87gHOTqeMYeWUsKaAbHIfcJNesEuDLxyKUHx3zx5/kuGbiU0+JUMzM89ziXDZkep/BWrWteVYEc0EXB6P7dFJCaYpR+YvUp4HE+z/bdlL/Q7TD8c+QsZ5aOoUI8wL69bQS2svzwlnozWorimGhfXCX23ERG9ZaTqHwIDAQAB" // 可选，支付宝提供给我们用于签名验证的公钥，通过支付宝管理后台获取
const privateKey = `MIIEowIBAAKCAQEAzlc+AaeVY8T1MwJofUfVK6eZdlHguaxJxyy3/gW1eTzwpBmc
619LSMKMO/AXePXgTKTIAcrb1xi47RTWlXPNKLAdNaIve/rYTxYvH4xjTEjuhOWN
a0/r2KfKt8RTK0xdHZFo8A0z2PZiKbX08QVIZAgG0IRj3RhvJNSF4xoaSmfqxrWQ
40yaDSx1++Bx44nmBJfrdumbpiTXNeRNFXVTlLosaU+QNeko2Mw82RJ6UoIxpHdA
u3xmm9Y/S00F/rAQJIgbCQsr8NEtYLnjZaYHDcjHR59E5LRadRy3DEQD4JNBx27i
9OXkYU85i4tPUOyIejp7jmQk6xpADubgCJk/XwIDAQABAoIBAG7JwO06gaeT9ONy
g4/gpcBOuMMiqzqGGwbqJ9AoWIvEAKDbb7mg7NcgPhNgkfnMaqY8Q3dS6D7Rc1k4
Ow77okgaZ1SIxF4ZDVNJYfvacOZgslCAwDHCM6ucS+fnbZBt7AEYNfPN4uy3PXWP
0PgoEfpDpQUIUGZB3Es2IL0TItKXiVRHDIK3CLqXLp9CqCa6zRU7Sedq1xw9yDK5
kg6gCKMG7yA6tNoCUIfQbJepyCaBa6e7iDrqrVorU1OV9QdJNkmzbfHe2W8Oaexp
HgePi3HxNeo812bT2A+Xd0RkXv3sm42E6IbUyLxQC0b+U2rDulwtEzWD/ZlfbVKr
kMkE6OECgYEA+32Z1fviwzPAi4nt+cLQ+18B08XR6OxmqZBQfz/MV3lk17W8FznD
5PvsAKHo73v1T9zYAuOaXcG+Pwaiv7qlTTAVWOe8+dZHsAjoSaVgGaJ4Sw03UoXu
5giIn2XX24YDSeLLIkt8QNDRNP7fSAiuZdfjR6NGp0VMJz1p0JJdUCsCgYEA0gpk
rh2pIzonQhQ7sPOfqI8AWJEiav+qamfXdTEqwDHrLznGmErPqTwOyyh6ajbGfGJT
9T99ZuZPqPlBa3KWSd/NpI3f94KD9A3LEHLLmEJk0+m37tn16m6edeihJ5wKyszr
S/CzPPJpLz9W5BdjfhlHo2n4trvMPnWjzC0tv50CgYEA69d+s4hSwIJA19rSe/2x
Y3pWSVXjIw7Gu88lXh/jLkkeQ5gfOpymU4/YY0NzIVra0zkkrZjqA3CUS6CacTjE
md00t7oioxzK/49q3t2igIe3TZKoRdBqF/j6vpiQVEKZxOlVW+T0MghhmNRHadMS
UVmiv2Sj9mtrHgVevQT1Ux8CgYAfjJga3zGsrOuArXFZ3v44X2J56zL0R3rdiMOA
QuZdACKHcXI3JeWTUgYW1fmtQpUF0e8yg1revfFxPB9reEbCye8lrHnbv6r6WagK
zHNwZ2ilvBF4uxyJHhyHMW4jePjWBXnCamFB4leSVT1D/Y9gPYPZM+275PGE4D6+
SHr6sQKBgA02odehe1yvh77rm+VOfesigPREB9zoqBT6z4WLzvdVzBerSWtr3OnF
ClBAUzxjWb8lg6D2jwTKkDYTKxmb0iMUdwFZQPmaXIZGQiFcq0zV+Khgykgz5cW2
HVX/ulqg/ZBadr7Od1eRX/DLhB/4wj12Y5BaVzavJJ01+8ptpKH3`
var client = alipay.New(appId, aliPublicKey, privateKey, true)

// TradeController operations for Trade
type TradeController struct {
	BaseController
}

// URLMapping ...
func (c *TradeController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Trade
// @Param	body		body 	models.Trade	true		"body for Trade content"
// @Success 201 {int} models.Trade
// @Failure 403 body is empty
// @router / [post]
func (c *TradeController) Post() {
	var v models.Trade
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddTrade(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
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
// @Description get Trade by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Trade
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TradeController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTradeById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Trade
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Trade
// @Failure 403
// @router / [get]
func (c *TradeController) GetAll() {
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

	l, err := models.GetAllTrade(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Trade
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Trade	true		"body for Trade content"
// @Success 200 {object} models.Trade
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TradeController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Trade{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateTradeById(&v); err == nil {
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
// @Description delete the Trade
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TradeController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTrade(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @router /pay
func (c *TradeController) Pay(){
	var p = alipay.AliPayTradePagePay{}
	host := GetHostName()
	p.NotifyURL = fmt.Sprintf("%v/v1/trade/alipay_notify", host)
	p.ReturnURL = fmt.Sprintf("%v/v1/trade/alipay_return", host)

	user, err := c.GetUser()
	u1, err := uuid.NewV4()
	if err != nil{
		c.Data["json"] = map[string]interface{}{"success": false, "msg": err.Error()}
		c.ServeJSON()
		return
	}

	uid := fmt.Sprintf("%s", u1)
	trade := models.Trade{UserId: user.Id, TradeNo: uid}

	o := orm.NewOrm()
	_, err = o.Insert(&trade)

	if err != nil{
		c.Data["json"] = map[string]interface{}{"success": false, "msg": err.Error()}
		c.ServeJSON()
		return
	}

	p.Subject = "高考题库"
	p.OutTradeNo = uid
	p.TotalAmount = "0.01"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}

	var payURL = url.String()
	fmt.Println("url", payURL)
	c.Data["json"] = map[string]interface{}{"success": true, "url": payURL}
	c.ServeJSON()
}


// @router /alipay_return
func (c *TradeController) Return() {
	c.Ctx.Request.ParseForm()
	form := c.Ctx.Request.Form
	ok, err := client.VerifySign(form)
	msg := ""
	fmt.Println(ok, err)
	if ok{
		out_trade_no := form["out_trade_no"]
		fmt.Println("form: ", form)
		if len(out_trade_no) > 0{
			fmt.Println("out_trade_no: ", out_trade_no[0])
			err := markRecordPaySuccess(out_trade_no[0])
			if err != nil{
				msg = fmt.Sprintf("标记%v失败：%v", out_trade_no, err.Error())
			}else{
				msg = "支付成功"
			}
		}else{
			msg = fmt.Sprintf("支付成功，但获取订单号失败")
		}

	}else{
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("支付失败: %v", err.Error())))
	}
	c.Ctx.ResponseWriter.Write([]byte(msg))
}


func markRecordPaySuccess(s string) error {
	fmt.Println("mark ", s)
	o := orm.NewOrm()
	qs := o.QueryTable("trade")
	trade := models.Trade{}
	err := qs.Filter("trade_no", s).One(&trade)

	if err != nil{
		return err
	}

	o.Read(&trade)
	fmt.Println(trade.Active)
	if trade.Active != 1{
		trade.Active = 1
		o.Update(&trade)
		user, err := models.GetUsersById(trade.UserId)
		if err == nil{
			fmt.Println("copy record")
			go models.CopyAnkiDeckToMemPlus(user)
			return err
		}else{
			return err
		}
	}

	return nil
}

// Post ...
// @router /alipay_notify [post]
func (c *TradeController) Notify() {
	var noti, err = client.GetTradeNotification(c.Ctx.Request)
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	var msg string
	if noti != nil {
		err = markRecordPaySuccess(noti.OutTradeNo)
		if err != nil{
			msg = err.Error()
		}else{
			msg = "支付成功"
		}

	} else {
		msg = "支付失败"
	}
	fmt.Println(msg)
	c.Ctx.ResponseWriter.Write([]byte(msg))
}