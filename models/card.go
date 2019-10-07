package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Card struct {
	Id             int       `orm:"column(id);auto"`
	CreatedAt      time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt      time.Time `orm:"column(updated_at);type(timestamp);null"`
	DeletedAt      time.Time `orm:"column(deleted_at);type(timestamp);null"`
	Status         string    `orm:"column(status);size(255);null"`
	Level          int       `orm:"column(level);null"`
	NextTrigger    time.Time `orm:"column(trigger_start_time);type(timestamp);null"`
	Note           *Note     `orm:"rel(one);column(nid);null"`
	Loop           *Loop     `orm:"rel(one);column(loop_id);null"`
	Finish         bool      `orm:"column(finish);null"`
	Did            int       `orm:"column(did);null"`
	Uid            int       `orm:"column(uid);null"`
}

func (t *Card) TableName() string {
	return "card"
}

func init() {
	orm.RegisterModel(new(Card))
}

func NewCard() Card {
	c := Card{}
	c.Loop = &Loop{Id: 1}
	return c
}

// AddCard insert a new Card into database and returns
// last inserted Id on success.
func AddCard(m *Card) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCardById retrieves Card by Id. Returns error if
// Id doesn't exist
func GetCardById(id int) (v *Card, err error) {
	o := orm.NewOrm()
	v = &Card{Id: id}
	if err = o.Read(v); err == nil {
		o.LoadRelated(v, "nid")
		return v, nil
	}
	return nil, err
}

// GetAllCard retrieves all Card matches certain condition. Returns empty list if
// no records exist
func GetAllCard(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Card))
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

	var l []Card
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

// UpdateCard updates Card by Id and returns error if
// the record to be updated doesn't exist
func UpdateCardById(m *Card) (err error) {
	o := orm.NewOrm()
	v := Card{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCard deletes Card by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCard(id int) (err error) {
	o := orm.NewOrm()
	v := Card{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Card{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}


func GetNextTriggerTime(this *Card) (nextTriggerTime time.Time, err error, Finish bool){
	now := time.Now()
	o := orm.NewOrm()
	_, err = o.LoadRelated(this, "loop_id")
	if err != nil{
		return nextTriggerTime, errors.New("load loop id fail"), false
	}

	loopDurs := this.Loop.GetLoops()
	lev := this.Level
	if lev < 0{
		lev = 0
	}
	if len(loopDurs) <= this.Level{
		err = errors.New("CardComplete")
	}else{
		nextTriggerTime = now.Add(loopDurs[lev])
	}
	return
}


func RememberCard(id int) (card *Card, err error) {
	o := orm.NewOrm()
	card = &Card{Id: id}
	if err = o.Read(card); err == nil {
		now := time.Now()

		if card.NextTrigger.Unix() > now.Unix(){
			return card, errors.New("还未到复习时间")
		}

		card.Level ++
		nextTrigger, err, complete := GetNextTriggerTime(card)
		if err != nil{
			return card, errors.New(fmt.Sprintf("%v: %v", "获取复习时间失败", err))
		}
		card.Finish = complete
		card.NextTrigger = nextTrigger
		_, err = o.Update(card)
		if err == nil{
			LogCardHistory(card, "complete")
		}else{
			return card, errors.New(fmt.Sprintf("%v: %v", "更新失败", err.Error()))
		}
	}
	return
}


func ForgetCard(id int) (task *Card, err error) {
	o := orm.NewOrm()
	task = &Card{Id: id}
	if err = o.Read(task); err == nil {
		now := time.Now()
		if task.NextTrigger.Unix() > now.Unix(){
			err = errors.New("还未到复习时间")
			return
		}
		task.Level--
		if(task.Level < 0 ){
			task.Level = 0
		}
		task.NextTrigger = now
		_, err = o.Update(task)
		if err == nil{
			// LogCardHistory(task, "complete")
		}
	}
	return
}


func SosoCard(id int) (task *Card, err error) {
	o := orm.NewOrm()
	task = &Card{Id: id}
	if err = o.Read(task); err == nil {
		now := time.Now()
		if task.NextTrigger.Unix() > now.Unix(){
			err = errors.New("还未到复习时间")
			return
		}

		nextTrigger, err, finish := GetNextTriggerTime(task)
		if err != nil{
			return nil, err
		}
		task.Finish = finish
		task.NextTrigger = nextTrigger
		o.Update(task)
		_, err = o.Update(task)
		if err == nil{
			LogCardHistory(task, "complete")
		}
	}
	return
}

func IsReadyCard(this *Card)bool{
	now := time.Now()
	return this.NextTrigger.Unix() > now.Unix()
}

func IsNewCard(this *Card)bool{
	return this.Level == 0
}
