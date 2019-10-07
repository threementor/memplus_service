package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"time"
)

type Deck struct {
	Id        int    `orm:"column(id);auto"`
	Title     string `orm:"column(title);size(200);null"`
	ParentId  int    `orm:"column(parent_id);null"`
	AllCardCount int
	OwnCardCount int
}

func (t *Deck) TableName() string {
	return "deck"
}

func init() {
	orm.RegisterModel(new(Deck))
}

// AddDeck insert a new Deck into database and returns
// last inserted Id on success.
func AddDeck(m *Deck) (id int64, err error) {
	o := orm.NewOrm()
	if m.Title == ""{
		return 0, errors.New("名称不能为空")
	}
	id, err = o.Insert(m)
	return
}

// GetDeckById retrieves Deck by Id. Returns error if
// Id doesn't exist
func GetDeckById(id int) (v *Deck, err error) {
	o := orm.NewOrm()
	v = &Deck{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDeck retrieves all Deck matches certain condition. Returns empty list if
// no records exist
func GetAllDeck(query map[string]string, fields []string, sortby []string, order []string,
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
func UpdateDeckById(m *Deck) (err error) {
	o := orm.NewOrm()
	v := Deck{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		o.Begin()
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
		hasCycle, err := CycleCheck(&v, o)
		if err != nil{
			return err
		}
		if hasCycle{
			o.Rollback()
			return errors.New("牌组有环")
		}else{
			o.Commit()
		}
	}
	return
}

// DeleteDeck deletes Deck by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDeck(id int) (err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("deck")
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

func CycleCheck(this *Deck, o orm.Ormer) (bool, error) {
	childs := GetSons(this, o)
	hasShow := map[int]int {this.Id: 1}

	for len(childs) > 0{
		pop := childs[0]
		childs = childs[1:]
		if hasShow[pop.Id] == 1{
			return true, nil
		}else{
			hasShow[pop.Id] = 1
		}
		newChilds := GetSons(pop, o)
		childs = append(childs, newChilds...)
	}
	return false, nil
}

func GetSons(this *Deck, o orm.Ormer)[]*Deck{
	sons := []*Deck{}
	qs := o.QueryTable("deck")
	qs.Filter("parent_id", this.Id).All(&sons)
	return sons
}

func GetSubDirs(this *Deck, with_self bool)([]*Deck, error){
	if this == nil{
		return nil, errors.New("it's null dir")
	}
	childs := []*Deck{}
	o := orm.NewOrm()
	qs := o.QueryTable("deck")
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

func GetSonCard(this *Deck)([]*Card){
	cards := []*Card{}
	o := orm.NewOrm()
	qs := o.QueryTable("card")
	qs.Filter("did", this.Id).All(&cards)
	return cards
}


func GetCards(this *Deck, user *User)([]*Card, error){
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
	qs.Filter("uid", user.Id).Filter("did__in", sub_dir_ids).All(&cards)
	return cards, nil
}


func syncCards(this *Deck, user *User)(error) {
	//1 生成左右两份 map
	//2 将map中一样的部分消除
	//3 将nid表多出来的，插入card
	//4 将card表多出来的，删除

	nidMap := map[int]int{}
	cidMap := map[int]*Card{}
	o := orm.NewOrm()

	qs1 := o.QueryTable("note")
	notes := []*Note{}
	qs1.Filter("did", this.Id).All(&notes)
	for i:=0; i<len(notes); i++{
		n := notes[i]
		nidMap[n.Id] = 1
	}

	qs2 := o.QueryTable("card")
	cards := []*Card{}
	qs2.Filter("uid", user.Id).Filter("did", this.Id).All(&cards)
	for i:=0; i<len(cards); i++{
		cidMap[cards[i].Note.Id] = cards[i]
	}

	// 都有的 不去管
	// cid 没有的，添加进去

	allHave := map[int]int{}
	cidLess := map[int]int{}
	for _, note := range notes {
		nid := note.Id
		if cidMap[nid] != nil{
			allHave[nid] = 1
		}else{
			beego.Info("have no", nid)
			cidLess[nid] = 1
		}
	}

	for nid, _ := range allHave{
		delete(cidMap, nid)
		delete(nidMap, nid)
	}
	// cid有，nid没有的，删掉cid
	for _, card := range cidMap{
		o.Delete(card)
	}
	// nid有，cid没有的，添加到user名下
	for nid, _ := range cidLess{
		beego.Info("insert %v", nid)
		c := Card{Note: &Note{Id: nid}, Uid: user.Id, Did: this.Id, Level: 0, Loop: &Loop{Id: 1}}
		o.Insert(&c)
	}
	return nil
}


func GetReadyCards(this *Deck, user *User)([]*Card, error){
	//筛选子目录下的所有card，然后过滤readytime
	err := syncCards(this, user)
	if err != nil{
		return nil, err
	}
	cards, err := GetCards(this, user)
	now := time.Now()
	ready_cards := []*Card{}
	if err != nil{
		return ready_cards, err
	}
	o := orm.NewOrm()
	// 筛选出来没有建立task的进行新建
	for i:=0; i<len(cards); i++{
		card := cards[i]
		if card.NextTrigger.Unix() < now.Unix(){
			ready_cards = append(ready_cards, card)
		}
	}

	if len(ready_cards) > 100{
		ready_cards = ready_cards[:100]
	}
	//relation cardInfo
	for i:=0; i<len(ready_cards); i++{
		card := ready_cards[i]
		o.LoadRelated(card, "note")
	}
	return ready_cards, nil
}

func buildMemCardFromAnkiCard(ankiCard *AnkiCard, o orm.Ormer) (*Card, error){
	note := Note{Title: ankiCard.Q, Content: ankiCard.A, Type: "anki"}
	nid, err := o.Insert(&note)
	if err != nil{
		return nil, errors.New(fmt.Sprintf("insert note error, %v", err.Error()))
	}
	note.Id = int(nid)
	newCard := Card{Note: &note, Loop: &Loop{Id: 1}}
	_, e := o.Insert(&newCard)
	if e != nil{
		return nil, errors.New(fmt.Sprintf("insert card fail, %v", e.Error()))
	}
	return &newCard, nil
}

func copyCards(deck *AnkiDeck, newDir *Deck, user *User, o orm.Ormer) error{
	ankiCards := []*AnkiCard{}
	qs := o.QueryTable("anki_card")
	qs.Filter("did", deck.DeckId).All(&ankiCards)

	//批量插入notes
	notes := []*Note{}
	for i:=0; i<len(ankiCards); i++{
		dc := ankiCards[i]
		note := Note{Title: dc.Q, Content: dc.A, Type: "anki", Did: newDir.Id}
		notes = append(notes, &note)
	}
	o.InsertMulti(len(notes), notes)

	//根据notes批量插入cards
	sql := fmt.Sprintf("insert into card (level, nid, did, finish, loop_id) select 0, id, did, 0, 1 from note where did=%v", newDir.Id)
	_, err := o.Raw(sql).Exec()
	return err
}


func GetRootDirs(user *User) ([]*Deck, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("deck")
	dirs := []*Deck{}
	_, err := qs.Filter("user_id", user.Id).All(&dirs)
	return dirs, err
}


func CopyAnkiDeckToMemPlus(trade *Trade, user *User) error {
	//先copy新deck
	//批量插入note
	//批量生成cards

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
		newDir := &Deck{Title: name}
		did, err := o.Insert(newDir)

		newDir.Id = int(did)
		rela := UserDeckRela{Uid: user.Id, Deck: &Deck{Id: newDir.Id}}
		_, err = o.Insert(&rela)

		if err != nil{
			return err
		}
		newDirs[deck.DeckId] = newDir
	}

	success := true
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
			success = false
		}
	}
	if success{
		trade.Status = "finish"
	}else{
		trade.Status = "fail"
	}
	o.Update(trade)
	o.Commit()
	return nil
}

func refreshDeck(rela *UserDeckRela, o orm.Ormer, handles map[int]*UserDeckRela) (*UserDeckRela, error){

	if d, ok := handles[rela.Id]; ok{
		return d, nil
	}
	o.LoadRelated(rela, "did")
	qs := o.QueryTable("card").Filter("did", rela.Deck.Id)
	cardsCount, err := qs.Count()
	if err != nil{
		return nil, err
	}

	newCount, err := qs.Filter("level", 0).Count()
	if err != nil{
		return nil, err
	}
	now := time.Now()
	readyCount, err := qs.Filter("trigger_start_time__lt", now).Count()
	if err != nil{
		return nil, err
	}

	rela.ReadyCount = int(readyCount)
	rela.NewCount = int(newCount)
	rela.Deck.OwnCardCount = int(cardsCount)
	rela.Deck.AllCardCount = int(cardsCount)

	son_decks := GetSons(rela.Deck, o)
	son_deck_ids := []int{}
	for i:=0; i<len(son_decks); i++{
		dc := son_decks[i]
		son_deck_ids = append(son_deck_ids, dc.Id)
	}
	son_rela_deck := []*UserDeckRela{}
	if len(son_deck_ids) > 0{
		o.QueryTable(TABLE_USER_DECK_RELA).Filter("did__in", son_deck_ids).Filter("uid", rela.Uid).All(&son_rela_deck)
		for i:=0; i<len(son_rela_deck); i++{
			son := son_rela_deck[i]
			son, err = refreshDeck(son, o, handles)
			if err != nil{
				return nil, err
			}
			rela.ReadyCount += son.ReadyCount
			rela.NewCount += son.NewCount
			rela.Deck.AllCardCount += son.Deck.AllCardCount
			fmt.Println(rela.Deck.Title, son.Deck.Title, son.Deck.AllCardCount)
		}
	}

	o.Update(rela)
	o.Update(rela.Deck)
	handles[rela.Id] = rela
	return rela, nil
}

func RefreshCount(relas []*UserDeckRela) {
	o := orm.NewOrm()
	handles := map[int]*UserDeckRela{}
	for i:=0; i<len(relas); i++{
		_, err := refreshDeck(relas[i], o, handles)
		if err != nil{
			fmt.Println(err.Error())
		}

	}
	return
}


func GetDeckForUser(user *User)(decks []*Deck, err error){
	o := orm.NewOrm()
	relas := []*UserDeckRela{}
	qs := o.QueryTable("user_deck_rela")
	qs.Filter("uid", user.Id).All(&relas)
	if err != nil{
		return nil, err
	}
	ids := []int{}
	for i:=0; i<len(relas); i++{
		ids = append(ids, relas[i].Deck.Id)
	}
	if len(ids) != 0{
		qs = o.QueryTable("deck")
		qs.Filter("id__in", ids).All(&decks)
	}
	return decks, nil

}