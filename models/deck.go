package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"github.com/astaxie/beego/orm"
)

type Deck struct {
	Id        int    `orm:"column(id);auto"`
	Title     string `orm:"column(title);size(200);null"`
	ParentId  int    `orm:"column(parent_id);null"`
	UserId int    `orm:"column(user_id);null"`
}

func (t *Deck) TableName() string {
	return "deck"
}

func init() {
	orm.RegisterModel(new(Deck))
}

// AddKlgDir insert a new Deck into database and returns
// last inserted Id on success.
func AddKlgDir(m *Deck) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetKlgDirById retrieves Deck by Id. Returns error if
// Id doesn't exist
func GetKlgDirById(id int) (v *Deck, err error) {
	o := orm.NewOrm()
	v = &Deck{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllKlgDir retrieves all Deck matches certain condition. Returns empty list if
// no records exist
func GetAllKlgDir(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Deck))
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

	var l []Deck
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

// UpdateKlgDir updates Deck by Id and returns error if
// the record to be updated doesn't exist
func UpdateKlgDirById(m *Deck) (err error) {
	o := orm.NewOrm()
	v := Deck{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteKlgDir deletes Deck by Id and returns error if
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

	v := Deck{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Deck{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetSubDirs(this *Deck, with_self bool)([]*Deck, error){
	if this == nil{
		return nil, errors.New("it's null dir")
	}
	childs := []*Deck{}
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


func GetCards(this *Deck)([]*Card, error){
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
		sub_dir_ids = append(sub_dir_ids, sd.Id)
	}
	qs.Filter("klg_dir_id__in", sub_dir_ids).All(&cards)
	return cards, nil
}


func GetReadyCards(this *Deck, user *User)([]*Card, error){
	//筛选子目录下的所有card，然后过滤readytime
	cards, err := GetCards(this)
	now := time.Now()
	if err != nil{
		return nil, err
	}
	o := orm.NewOrm()
	ready_cards := []*Card{}
	// 筛选出来没有建立task的进行新建
	for i:=0; i<len(cards); i++{
		card := cards[i]
		if card.TriggerStartTime < int(now.Unix()){
			ready_cards = append(ready_cards, card)
		}
	}
	//relation cardInfo
	for i:=0; i<len(ready_cards); i++{
		card := ready_cards[i]
		o.LoadRelated(card, "note")
	}
	return ready_cards[:100], nil
}

func buildMemCardFromAnkiCard(ankiCard *AnkiCard) (*Card, error){
	o := orm.NewOrm()
	note := Note{Title: ankiCard.Q, Content: ankiCard.A, Type: "anki"}
	nid, err := o.Insert(note)
	if err != nil{
		return nil, errors.New(fmt.Sprintf("insert note error, %v", err.Error()))
	}
	note.Id = int(nid)
	newCard := Card{Note: &note}
	_, e := o.Insert(&newCard)
	if e != nil{
		return nil, errors.New(fmt.Sprintf("insert card fail, %v", e.Error()))
	}
	return &newCard, nil
}

func copyCards(deck *AnkiDeck, newDir *Deck, user *User, o orm.Ormer) error{
	//1 拿出来所有的anki card ， anki deck
	//2 根据anki card制作出 mem card
	//3. 根据 anki deck 制作出 mem deck
	ankiCards := []*AnkiCard{}
	qs := o.QueryTable("anki_card")
	qs.Filter("did", deck.DeckId).All(&ankiCards)

	for i:=0; i<len(ankiCards); i++{
		dc := ankiCards[i]
		_, err := buildMemCardFromAnkiCard(dc)
		if err != nil{
			return err
		}
	}
	return nil
}


func GetRootDirs(user *User) ([]*Deck, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("klg_dir")
	dirs := []*Deck{}
	_, err := qs.Filter("user_id", user.Id).All(&dirs)
	return dirs, err
}


func CopyAnkiDeckToMemPlus(user *User) error {
	o := orm.NewOrm()
	o.Begin()
	qs := o.QueryTable("anki_deck")
	var decks []*AnkiDeck
	qs.All(&decks)
	deckMap := map[int64]*AnkiDeck{}
	for i:=0; i<len(decks); i++{
		deckMap[decks[i].DeckId] = decks[i]
	}

	newDirs := map[int64]*Deck{}

	//初始化deck
	for _, deck := range deckMap{
		names := strings.Split(deck.Name, ":")
		name := deck.Name
		if len(names) > 1{
			name = names[len(names)-1]
		}
		newDir := &Deck{Title: name, UserId: user.Id}
		_, err := o.Insert(newDir)
		if err != nil{
			return err
		}
		newDirs[deck.DeckId] = newDir
	}
	//copy deck
	for deckId, newDir := range newDirs{
		deck := deckMap[deckId]
		if deck.Pdid != 0{
			parentDir := newDirs[deck.Pdid]
			newDir.ParentId = parentDir.Id
			o.Update(newDir)
		}
		err := copyCards(deck, newDir, user, o)
		if err != nil{
			return nil
		}
	}
	return nil
}
