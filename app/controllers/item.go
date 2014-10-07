package controllers

import (
	"github.com/revel/revel"
	"github.com/ericbutera/gofeeder/app/models"
)

type Item struct {
	*revel.Controller
}

func (c Item) Index(feedId int) revel.Result {
	if feedId > 0 {
		var items []models.Item
		Db.Raw(`
  SELECT
    i.*
  FROM items i
  INNER JOIN feed_items fi ON i.id = fi.item_id
  WHERE fi.feed_id = ?`, feedId).Scan(&items)
		return c.RenderJson(items)
	} else {
		return c.RenderJson(struct{ Error string }{"Invalid feedId"})
	}
}
