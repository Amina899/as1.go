package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type News struct {
	ID       int
	Title    string
	Content  string
	Author   string
	Created  time.Time
	Category string
}
