package controllers

import (
	"github.com/jgraham909/revmgo"
	"github.com/revel/revel"
	"labix.org/v2/mgo"
)

type App struct {
	*revel.Controller
	revmgo.MongoController
}

func (app *App) DB() *mgo.Database {
	dbName := revel.Config.StringDefault("mgo.database", "music_link_dev")
	return app.MongoSession.DB(dbName)
}

func (app *App) C(name string) *mgo.Collection {
	return app.DB().C(name)
}
