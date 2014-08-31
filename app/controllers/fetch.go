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
	feed := models.Feed{}
	Db.First(&feed)
	revel.INFO.Printf("Feed %+v", feed)

	// feeds := pull from db
	// for _, url := range feeds {
	fetched, err := rss.Fetch(feed.Url)
	if err != nil {
		revel.INFO.Printf("Unable to fetch %s: %s", feed.Url, err)
	} else {
		//feed.FetchedAt = time.Now()
		//Db.Update(&feed)
		Db.Model(&feed).UpdateColumn("FetchedAt", time.Now())
		revel.INFO.Printf("Feed %+v", feed)
		return c.RenderJson(fetched)
	}
	//handleFeed(fetched)
	// }

	return c.RenderJson(struct{ Moo string }{"HAHA"})
}

//func handleFeed(*) {
//}
