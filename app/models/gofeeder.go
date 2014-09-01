package models

import (
	"time"
)

type Feed struct {
	Id          int64
	Name        string    `sql:"type:text"`
	Url         string    `sql:"type:text"`
	DateFetched time.Time //ill handle this one
	CreatedAt   time.Time
	//UpdatedAt time.Time
	// DeletedAt time.Time //nothing but problems with this:  pq: invalid input syntax for type timestamp with time zone: "-0001-12-31T18:27:49-05:32:11"
}

type Item struct {
	Id      int64
	Name    string `sql:"type:text"`
	Url     string `sql:"type:text"`
	Content string `sql:"type:text"`
	// maybe add a content sha as an additional check?
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt time.Time
	Feeds []Feed `gorm:"many2many:feed_items;"`
}

/*
// there has to be a better way to do this, but meh
type FeedItems struct {
	FeedId int64
	ItemId int64
}
*/
