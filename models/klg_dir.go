package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"github.com/astaxie/beego/orm"
)

type KlgDir struct {
	Id        int    `orm:"column(id);auto"`
	Title     string `orm:"column(title);size(200);null"`
	ParentId  int    `orm:"column(parent_id);null"`
	UserId int    `orm:"column(user_id);null"`
}

func (t *KlgDir) TableName() string {
	return "klg_dir"
}

func init() {
	orm.RegisterModel(new(KlgDir))
}

// AddKlgDir insert a new KlgDir into database and returns
// last inserted Id on success.
func AddKlgDir(m *KlgDir) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetKlgDirById retrieves KlgDir by Id. Returns error if
// Id doesn't exist
func GetKlgDirById(id int) (v *KlgDir, err error) {
	o := orm.NewOrm()
	v = &KlgDir{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllKlgDir retrieves all KlgDir matches certain condition. Returns empty list if
// no records exist
func GetAllKlgDir(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(KlgDir))
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

	var l []KlgDir
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

// UpdateKlgDir updates KlgDir by Id and returns error if
// the record to be updated doesn't exist
func UpdateKlgDirById(m *KlgDir) (err error) {
	o := orm.NewOrm()
	v := KlgDir{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteKlgDir deletes KlgDir by Id and returns error if
// the record to be deleted doesn't exist
func DeleteKlgDir(id int) (err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("klg_dir")
	child_count, err := qs.Filter("parent_id", id).Count()
	if err != nil{
		return err
	}
	if child_count > 0{
		return errors.New("请先将子目录移走")
	}

	v := KlgDir{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&KlgDir{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetSubDirs(this *KlgDir, with_self bool)([]*KlgDir, error){
	if this == nil{
		return nil, errors.New("it's null dir")
	}
	childs := []*KlgDir{}
	o := orm.NewOrm()
	qs := o.QueryTable("klg_dir")
	qs.Filter("parent_id", this.Id).All(&childs)

	for _, c := range(childs){
		child_sub_dirs, err := GetSubDirs(c, false)
		if err != nil{
			return nil, err
		}
		childs = append(childs, child_sub_dirs...)
	}
	if with_self{
		println("with self", this.Id)
		childs = append(childs, this)
	}
	return childs, nil
}


func GetCards(this *KlgDir)([]*Card, error){
	//先获取所有子目录的id
	subDirs, err := GetSubDirs(this, true)

	if err != nil{
		return nil, err
	}
	//获取所有task, 并且ready_time > now的
	cards := []*Card{}
	o := orm.NewOrm()
	qs := o.QueryTable("card")
	sub_dir_ids := []int{}
	for _, sd := range(subDirs){
		println(sd.Id)
		sub_dir_ids = append(sub_dir_ids, sd.Id)
	}
	qs.Filter("klg_dir_id__in", sub_dir_ids).All(&cards)
	return cards, nil
}

func createNewTaskForCards(cardMap map[int]*Card, user *User) ( []*Task, error){
	// 没有的卡片重新建
	tasks := []*Task{}
	o := orm.NewOrm()
	for k, card := range cardMap{
		fmt.Println("build new", k)
		newTask := &Task{Card: card, Loop: &Loop{Id: 1}, UserId: user.Id}
		id, err := o.Insert(newTask)
		if err != nil{
			return tasks, err
		}else{
			newTask.Id = int(id)
			tasks = append(tasks, newTask)
		}
	}
	return tasks, nil
}


func GetReadyTasks(this *KlgDir, user *User)([]*Task, error){

	//先获取所有子目录的id
	cards, err := GetCards(this)
	now := time.Now()
	if err != nil{
		return nil, err
	}
	o := orm.NewOrm()
	qs := o.QueryTable("task")


	tasks := []*Task{}
	oldTasks := []*Task{}
	cardMap := map[int]*Card{}
	cardIds := []int{}
	hasNoTaskCards := map[int]*Card{}

	for _, card := range(cards){
		cardMap[card.Id] = card
		hasNoTaskCards[card.Id] = card
		cardIds = append(cardIds, card.Id)
	}

	// 筛选出来没有建立task的进行新建
	qs.Filter("card_id__in", cardIds).All(&tasks)
	for i:=0; i<len(tasks); i++{
		task := tasks[i]
		delete(hasNoTaskCards, task.Card.Id)
		if task.TriggerStartTime < int(now.Unix()){
			oldTasks = append(oldTasks, task)
		}
	}

	newTasks, err := createNewTaskForCards(hasNoTaskCards, user)
	if err != nil{
		return newTasks, err
	}
	newTasks = append(newTasks, oldTasks...)

	//relation cardInfo
	for i:=0; i<len(newTasks); i++{
		task := newTasks[i]
		task.Card = cardMap[task.Card.Id]
	}

	return newTasks[:100], nil
}


func copyCards(deck *AnkiDeck, newDir *KlgDir, user *User) error{
	deckCards := []*AnkiCard{}
	o := orm.NewOrm()
	qs := o.QueryTable("anki_card")
	qs.Filter("did", deck.DeckId).All(&deckCards)

	for i:=0; i<len(deckCards); i++{
		dc := deckCards[i]
		newCard := Card{Title: dc.Q, Content: dc.A, KlgDirId:newDir.Id, Type: "anki"}
		_, e := o.Insert(&newCard)
		if e != nil{
			return errors.New(fmt.Sprintf("insert fail, %v", e.Error()))
		}
	}
	return nil
}


func GetRootDirs(user *User) ([]*KlgDir, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("klg_dir")
	dirs := []*KlgDir{}
	_, err := qs.Filter("user_id", user.Id).All(&dirs)
	return dirs, err
}


func CopyAnkiDeckToMemPlus(user *User) error {
	o := orm.NewOrm()
	qs := o.QueryTable("anki_deck")
	var decks []*AnkiDeck
	println("2")
	qs.All(&decks)
	println("3")
	deckMap := map[int64]*AnkiDeck{}
	for i:=0; i<len(decks); i++{
		deckMap[decks[i].DeckId] = decks[i]
	}

	newDirs := map[int64]*KlgDir{}

	for _, deck := range deckMap{
		names := strings.Split(deck.Name, ":")
		name := deck.Name
		if len(names) > 1{
			name = names[len(names)-1]
		}
		newDir := &KlgDir{Title: name, UserId: user.Id}
		_, err := o.Insert(newDir)
		if err != nil{
			return err
		}
		newDirs[deck.DeckId] = newDir
	}

	for deckId, newDir := range newDirs{
		deck := deckMap[deckId]
		if deck.Pdid != 0{
			parentDir := newDirs[deck.Pdid]
			newDir.ParentId = parentDir.Id
			o.Update(newDir)
		}

		err := copyCards(deck, newDir, user)
		if err != nil{
			return nil
		}
	}
	return nil
}
