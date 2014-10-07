package controllers

import (
	"github.com/revel/revel"
	"github.com/ericbutera/gofeeder/app/models"
)

type Feed struct {
	*revel.Controller
}

func (c Feed) Index() revel.Result {
	var feeds []models.Feed
	Db.Find(&feeds)
	return c.RenderJson(feeds)
}
