package models

import (
	"time"
)

type Feed struct {
	Id        int64
	Name      string `sql:"type:text"`
	Url       string `sql:"type:text"`
	FetchedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Item struct {
	Id      int64
	Name    string `sql:"type:text"`
	Url     string `sql:"type:text"`
	Content string `sql:"type:text"`
	// maybe add a content sha as an additional check?
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
