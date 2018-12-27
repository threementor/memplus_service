package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"time"
)

type Trade struct {
	Id     int  `orm:"column(id);auto"`
	UserId int  `orm:"column(user_id);null"`
	Pay int `orm:"column(pay);null"`
	Status string `orm:"column(status);null"`
	TradeNo string `orm:"column(trade_no);null"`
	Product *Product `orm:"rel(one);column(product_id);null"`
	CreateTime time.Time `orm:"column(ct);type(timestamp);null"`
	PayTime time.Time `orm:"column(pt);type(timestamp);null"`
	Amount float32 `orm:"column(amount);null"`
}

//pay
var TRADE_PAY_PAID = 1
var TRADE_PAY_NO_PAID = 0
var TRADE_PAY_BACK = -1

var pay_map = map[int][]string{
	TRADE_PAY_NO_PAID: []string{"未付款", "unpay"},
	TRADE_PAY_PAID: []string{"已付款", "paid"},
	TRADE_PAY_BACK: []string{"已退款", "back"},
}

//status
var TRADE_STATUS_INITING = "initing"
var TRADE_STATUS_FINISH = "finish"


var trade_map = map[string]string{
	TRADE_STATUS_INITING: "发货中",
	TRADE_STATUS_FINISH: "已交付",
}


func (t *Trade) TableName() string {
	return "trade"
}

func (t *Trade) AsMap() map[string]interface{} {
	o := orm.NewOrm()
	o.LoadRelated(t, "product_id")
	m := map[string]interface{}{}
	m["Id"] = t.Id
	payInfo, ok := pay_map[t.Pay]
	if ok{
		m["pay"] = payInfo[1]
		m["payHuman"] = payInfo[0]
	}else{
		m["pay"] = t.Pay
		m["payHuman"] = t.Pay
	}

	m["status"] = t.Status
	tradeStatus, ok := trade_map[t.Status]
	if ok{
		m["statusHuman"] = tradeStatus
	}else{
		m["statusHuman"] = t.Status
	}

	m["tradeNo"] = t.TradeNo
	m["product"] = t.Product
	m["payTime"] = t.PayTime.Format("2006-01-02 15:04:05")
	m["createTime"] = t.CreateTime.Format("2006-01-02 15:04:05")
	m["amount"] = t.Amount
	return m
}

func init() {
	orm.RegisterModel(new(Trade))
}

// AddTrade insert a new Trade into database and returns
// last inserted Id on success.
func AddTrade(m *Trade) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTradeById retrieves Trade by Id. Returns error if
// Id doesn't exist
func GetTradeById(id int) (v *Trade, err error) {
	o := orm.NewOrm()
	v = &Trade{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTrade retrieves all Trade matches certain condition. Returns empty list if
// no records exist
func GetAllTrade(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Trade))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Trade
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTrade updates Trade by Id and returns error if
// the record to be updated doesn't exist
func UpdateTradeById(m *Trade) (err error) {
	o := orm.NewOrm()
	v := Trade{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTrade deletes Trade by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTrade(id int) (err error) {
	o := orm.NewOrm()
	v := Trade{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Trade{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

