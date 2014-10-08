package controllers

import (
	"github.com/ericbutera/gofeeder/app/models"
	"github.com/revel/revel"
)

type Item struct {
	*revel.Controller
}

type ListItem struct {
	Id   int64
	Name string
	Url  string
}

func (c Item) Index(feedId int) revel.Result {
	if feedId > 0 {
		var items []ListItem
		Db.Raw(`
  SELECT
    i.Id, i.Name, i.Url
  FROM items i
  INNER JOIN feed_items fi ON i.id = fi.item_id
  WHERE fi.feed_id = ?`, feedId).Scan(&items)
		return c.RenderJson(items)
	} else {
		return c.RenderJson(struct{ Error string }{"Invalid feedId"})
	}
}

func (c Item) View(id int) revel.Result {
	if id > 0 {
		item := models.Item{}
		Db.First(&item, id)
		return c.RenderJson(item)
	} else {
		return c.RenderJson(struct{ Error string }{"Invalid itemId"})
	}
}
