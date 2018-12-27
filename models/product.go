package models

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Product struct {
	Id     int     `orm:"column(id);auto"`
	Title  string  `orm:"column(title);size(200);null"`
	Price  float32 `orm:"column(price);null"`
	Status int8    `orm:"column(status);null"`
	Detail string  `orm:"column(detail);null"`
}

func (t *Product) TableName() string {
	return "product"
}

func init() {
	orm.RegisterModel(new(Product))
}

// AddProduct insert a new Product into database and returns
// last inserted Id on success.
func AddProduct(m *Product) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProductById retrieves Product by Id. Returns error if
// Id doesn't exist
func GetProductById(id int) (v *Product, err error) {
	o := orm.NewOrm()
	v = &Product{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProduct retrieves all Product matches certain condition. Returns empty list if
// no records exist
func GetAllProduct(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Product))
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

	var l []Product
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

// UpdateProduct updates Product by Id and returns error if
// the record to be updated doesn't exist
func UpdateProductById(m *Product) (err error) {
	o := orm.NewOrm()
	v := Product{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProduct deletes Product by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProduct(id int) (err error) {
	o := orm.NewOrm()
	v := Product{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Product{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func SubmitTradeForProduct(id int, user *User) (*Trade, error) {
	o := orm.NewOrm()
	product := Product{Id: id}
	// ascertain id exists in the database
	if err := o.Read(&product); err == nil {
		u1, err := uuid.NewV4()
		uid := fmt.Sprintf("%s", u1)
		now := time.Now()
		trade := Trade{Product: &product, UserId: user.Id, TradeNo: uid, CreateTime: now, Amount: product.Price}
		tid, err := o.Insert(&trade)
		if err != nil{
			return nil, err
		}
		trade.Id = int(tid)
		return &trade, nil
	}
	return nil, errors.New("产品不存在或已下架")
}