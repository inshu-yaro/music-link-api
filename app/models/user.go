package models

import (
	fb "github.com/huandu/facebook"
	"labix.org/v2/mgo/bson"
	"time"
)

type User struct {
	Id        bson.ObjectId `json:"id"         bson:"_id"`
	FirstName string        `json:"first_name" bson:"first_name"`
	LastName  string        `json:"last_name"  bson:"last_name"`
	Birthday  time.Time     `json:"birtday"    bson:"birthday"`
	Token     string        `json:"token"      bson:"token"`
	Email     string        `json:"email"      bson:"email"`
}

func (u *User) LoadFromFacebook() {

}
