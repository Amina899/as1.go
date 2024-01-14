package models

import (
	"database/sql"
	"time"
)

type NewsModel struct {
	DB *sql.DB
}

type News struct {
	ID       int
	Title    string
	Content  string
	Author   string
	Created  time.Time
	Category string
	Expires  time.Time
}

type NewsCategory struct {
	ID   int
	Name string
}

type NewsWithCategory struct {
	News
	Category NewsCategory
}
