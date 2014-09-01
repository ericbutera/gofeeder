// pull out into job/cli process
package controllers

import (
	"github.com/SlyMarbo/rss"
	"github.com/revel/revel"
	"gofeeder/app/models"
	"time"
)

type Fetch struct {
	*revel.Controller
}

func (c Fetch) Fetch() revel.Result {
	//feed := models.Feed{}
	//Db.First(&feed)
	//revel.INFO.Printf("Feed %+v", feed)

	var feeds []models.Feed
	Db.Find(&feeds)
	var count = 0
	for _, feed := range feeds {
		revel.INFO.Printf("Feed %s", feed.Name)
		fetched, err := rss.Fetch(feed.Url)
		if err != nil {
			revel.INFO.Printf("Unable to fetch %s: %s", feed.Url, err)
		} else {
			handleFeed(fetched, feed)
			count++

			Db.Model(&feed).UpdateColumn("DateFetched", time.Now())
			revel.INFO.Printf("Feed %+v", feed)
		}
	}

	ret := struct{ Count int }{count}
	return c.RenderJson(ret)
}

func handleFeed(rssFeed *rss.Feed, feed models.Feed) {
	for _, feedItem := range rssFeed.Items {
		item := models.Item{
			//Feeds:   []models.Feed{feed},
			Name:    feedItem.Title,
			Content: feedItem.Content,
			Url:     feedItem.Link}
		Db.Save(&item)

		sql := "INSERT INTO feed_items (item_id, feed_id) SELECT ?,? WHERE NOT EXISTS (SELECT * FROM feed_items WHERE item_id =? AND feed_id =?)"
		Db.Exec(sql, item.Id, feed.Id, item.Id, feed.Id)
	}
}
