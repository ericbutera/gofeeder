package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	/*feed := models.Feed{Name: "reddit", Url: "http://localhost:9000/public/reddit.xml"}
	Db.Create(&feed)
	revel.INFO.Printf("feed1 %+v", feed)
	Db.Save(&feed)*/

	/*feed2 := models.Feed{}
	Db.First(&feed2)
	revel.INFO.Printf("feed2 %+v", feed2)*/

	return c.Render()
}
