package controllers

import (
	"encoding/json"
	"memplus_service/models"
	"strconv"
)

// DeckController operations for Deck
type DeckController struct {
	LoginReqireController
}

// URLMapping ...
func (c *DeckController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Deck
// @Param	body		body 	models.Deck	true		"body for Deck content"
// @Success 201 {int} models.Deck
// @Failure 403 body is empty
// @router / [post]
func (c *DeckController) Post() {
	var v models.Deck
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		user, err := c.GetUser()
		if err != nil{
			c.Data["json"] = err.Error()
		}else{
			if did, err := models.AddDeck(&v); err == nil {
				rela := models.UserDeckRela{Uid: user.Id, Deck: &models.Deck{Id: int(did)}}
				if rid, err := models.AddUserDeckRela(&rela); err == nil{
					rela.Id = int(rid)
					c.SendSuccess(rela.AsMap())
				}else{
					c.SendError(err, -1)
				}
			} else {
				c.SendError(err, -1)
			}
		}


	} else {
		c.SendError(err, -1)
	}
}


// GetTasks ...
// @Title Get GetTasks
// @Description get GetTasks  by Deck id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Note
// @router /:id/ready_cards [get]
func (c *DeckController) GetReadyTasks(){
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	dir, err := models.GetDeckById(id)
	if err != nil{
		c.SendError(err, -1)
		return
	}
	user, _ := c.GetUser()
	cards, err := models.GetReadyCards(dir, user)
	if err != nil {
		c.SendError(err, -1)
		return
	} else {
		c.SendSuccess(cards)
		return
	}
	c.ServeJSON()
}


// GetOne ...
// @Title Get One
// @Description get Deck by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Deck
// @Failure 403 :id is empty
// @router /:id [get]
func (c *DeckController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	user, err := c.GetUser()
	if err != nil{
		c.SendError(err, -1)
		return
	}
	deck := &models.Deck{Id: id}
	v, err := models.GetUserDeckRela(user, deck)
	if err != nil{
		c.SendError(err, -1)
		return
	} else {
		c.SendSuccess(v.AsMap())
	}
}

// GetAll ...
// @Title Get All
// @Description get Deck
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Deck
// @Failure 403
// @router / [get]
func (c *DeckController) GetAll() {

	user, err := c.GetUser()
	if err == nil{
		decks, err := models.GetDeckForUser(user)
		if err == nil{
			c.SendSuccess(decks)
		}else{
			c.SendError(err, -1)
		}
	}else{
		c.SendError(err, -1)
	}
}


// GetAll ...
// @Title Get Root Dirs
// @Description get Deck
// @Success 200 {object} models.Deck
// @Failure 403
// @router /roots [get]
func (c *DeckController) GetRootDirs() {

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
// @Description update the Deck
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Deck	true		"body for Deck content"
// @Success 200 {object} models.Deck
// @Failure 403 :id is not int
// @router /:id [put]
func (c *DeckController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Deck{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateDeckById(&v); err == nil {
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
// @Description delete the Deck
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *DeckController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteDeck(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}


// Put ...
// @Title Put
// @Description update the Deck
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Deck	true		"body for Deck content"
// @Success 200 {object} models.Deck
// @Failure 403 :id is not int
// @router /:id/create/card [post]
func (c *DeckController) AddCardToDeck() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	_, err := models.GetDeckById(id)
	if err != nil{
		c.SendError(err, -1)
		return
	}
	var note models.Note
	var card models.Card
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &note); err == nil {
		note.Did = id
		if nid, err := models.AddNote(&note); err == nil {
			note.Id = int(nid)
			if err := json.Unmarshal(c.Ctx.Input.RequestBody, &card); err == nil {
				card.Loop = &models.Loop{Id: 1}
				card.Note = &note
				card.Did = id
				if _, err = models.AddCard(&card); err == nil {
					c.Ctx.Output.SetStatus(201)
					c.SendSuccess(nil)
					return
				} else {
					c.SendError(err, -1)
					return
				}
			}else{
				c.SendError(err, -1)
				return
			}
		}else{
			c.SendError(err, -1)
			return
		}
	} else {
		c.SendError(err, -1)
		return
	}
	c.SendSuccess(nil)
}