package controllers

import (
	"encoding/json"
	rev "github.com/revel/revel"
	"github.com/tuvistavie/music-link-api/app/models"
	"labix.org/v2/mgo/bson"
)

type Users struct {
	App
}

func (c *Users) Index() rev.Result {
	var users []models.User
	c.C("users").Find(bson.M{}).All(&users)
	return c.RenderJson(users)
}

func (c *Users) Create() rev.Result {
	var user models.User
	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		rev.ERROR.Println(err)
		return c.RenderJson(err)
	}
	if err := user.LoadFromFacebook(); err != nil {
		rev.ERROR.Println(err)
		return c.RenderJson(err)
	}
	user.Id = bson.NewObjectId()
	c.C("users").Insert(&user)
	return c.RenderJson(user)
}
