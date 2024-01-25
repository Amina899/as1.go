package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type News struct {
	ID      int
	Title   string
	Content string
	Author  string
	Created time.Time
}
type Category struct {
	Name string
}

type Contact struct {
	Name    string
	Email   string
	Message string
}
