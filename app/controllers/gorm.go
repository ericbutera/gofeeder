package controllers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
	"gofeeder/app/models"
)

var (
	Db *gorm.DB
)

func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Cound not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectionString() string {
	//host := getParamString("db.host", "")
	user := getParamString("db.user", "gofeeder")
	pass := getParamString("db.pass", "gofeeder")
	name := getParamString("db.name", "gofeeder")
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, pass, name)
}

func InitDB() {
	revel.INFO.Printf("inside init db!")
	//Db, err := gorm.Open("sqlite3", "/tmp/goreeder.db")
	db, err := gorm.Open("postgres", getConnectionString())
	if err != nil {
		revel.ERROR.Fatal(err)
	}
	db.DB()
	db.AutoMigrate(models.Feed{})
	db.AutoMigrate(models.Item{})
	//db.Create(models.Feed{Name: "reddit", Url: "http://localhost:9000/public/rss/reddit.xml"})
	//db.Create(models.Feed{Name: "hn", Url: "http://localhost:9000/public/rss/hn.xml"})
	//db.AutoMigrate(models.FeedItems{})

	// figure out how to test ENV
	db.LogMode(true)
	db.SetLogger(gorm.Logger{revel.INFO})

	//db.SingularTable(true)
	Db = &db
}

/*
type GorpController struct {
	*revel.Controller
	Db *gorm.DB
}
*/
