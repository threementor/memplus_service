package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Task struct {
	Id               int       `orm:"column(id);auto"`
	CreatedAt        time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt        time.Time `orm:"column(updated_at);type(timestamp);null"`
	DeletedAt        time.Time `orm:"column(deleted_at);type(timestamp);null"`
	Status           string    `orm:"column(status);size(255);null"`
	Level            int       `orm:"column(level);null"`
	TriggerStartTime int       `orm:"column(trigger_start_time);null"`
	TriggerDueTime   int       `orm:"column(trigger_due_time);null"`
	Card      *Card      `orm:"rel(one);column(card_id);null"`
	Loop           *Loop      `orm:"rel(one);column(loop_id);null"`
	Title            string    `orm:"column(title);size(255);null"`
	UserId           int      `orm:"column(user_id);null"`
	Finish           bool      `orm:"column(finish);null"`
}

func (t *Task) TableName() string {
	return "task"
}

func init() {
	orm.RegisterModel(new(Task))
}

// AddTask insert a new Task into database and returns
// last inserted Id on success.
func AddTask(m *Task) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTaskById retrieves Task by Id. Returns error if
// Id doesn't exist
func GetTaskById(id int) (v *Task, err error) {
	o := orm.NewOrm()
	v = &Task{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


// GetAllTask retrieves all Task matches certain condition. Returns empty list if
// no records exist
func GetAllTask(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Task))
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

	var l []Task
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

// UpdateTask updates Task by Id and returns error if
// the record to be updated doesn't exist
func UpdateTaskById(m *Task) (err error) {
	o := orm.NewOrm()
	v := Task{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTask deletes Task by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTask(id int) (err error) {
	o := orm.NewOrm()
	v := Task{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Task{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}



func GetNextTriggerTime(this *Task) (nextTriggerTime time.Time, err error, Finish bool){
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
		err = errors.New("TaskComplete")
	}else{
		nextTriggerTime = now.Add(loopDurs[lev])
	}
	return
}


func RememberTask(id int) (task *Task, err error) {
	o := orm.NewOrm()
	task = &Task{Id: id}
	if err = o.Read(task); err == nil {
		now := time.Now()

		if task.TriggerStartTime > int(now.Unix()){
			return task, errors.New("还未到复习时间")
		}

		task.Level ++
		nextTrigger, err, complete := GetNextTriggerTime(task)
		if err != nil{
			return task, errors.New(fmt.Sprintf("%v: %v", "获取复习时间失败", err))
		}
		task.Finish = complete
		task.TriggerStartTime = int(nextTrigger.Unix())
		task.TriggerDueTime = task.TriggerStartTime + 60 * 60 * 24
		_, err = o.Update(task)
		if err == nil{
			LogTaskHistory(task, "complete")
		}else{
			return task, errors.New(fmt.Sprintf("%v: %v", "更新失败", err.Error()))
		}
	}
	return
}


func ForgetTask(id int) (task *Task, err error) {
	o := orm.NewOrm()
	task = &Task{Id: id}
	if err = o.Read(task); err == nil {
		now := time.Now()
		if task.TriggerStartTime > int(now.Unix()){
			err = errors.New("还未到复习时间")
			return
		}
		task.Level--
		if(task.Level < 0 ){
			task.Level = 0
		}
		nextTrigger, err, finish := GetNextTriggerTime(task)
		if err != nil{
			return nil, err
		}
		task.TriggerStartTime = int(nextTrigger.Unix())
		task.TriggerDueTime = task.TriggerStartTime + 60 * 60 * 24
		task.Finish = finish
		_, err = o.Update(task)
		if err == nil{
			// LogTaskHistory(task, "complete")
		}
	}
	return
}


func SosoTask(id int) (task *Task, err error) {
	o := orm.NewOrm()
	task = &Task{Id: id}
	if err = o.Read(task); err == nil {
		now := time.Now()
		if task.TriggerStartTime > int(now.Unix()){
			err = errors.New("还未到复习时间")
			return
		}
	
		nextTrigger, err, finish := GetNextTriggerTime(task)
		if err != nil{
			return nil, err
		}
		task.Finish = finish
		task.TriggerStartTime = int(nextTrigger.Unix())
		task.TriggerDueTime = task.TriggerStartTime + 60 * 60 * 24
		o.Update(task)
		_, err = o.Update(task)
		if err == nil{
			LogTaskHistory(task, "complete")
		}
	}
	return
}
