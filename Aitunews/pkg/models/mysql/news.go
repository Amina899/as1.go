package mysql

import (
	"database/sql"
	"fmt"
	"log"
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

func NewNewsModel(db *sql.DB) *NewsModel {
	return &NewsModel{
		DB: db,
	}
}

func (m *NewsModel) InsertNews(title, content, author, category string, created, expires time.Time) (int, error) {
	stmt := `INSERT INTO news (title, content, author, category, created, expires) VALUES (?, ?, ?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, title, content, author, category, created, expires)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return int(id), nil
}

func (m *NewsModel) GetLatestNews(limit int) ([]News, error) {
	stmt := `SELECT id, title, content, author, created, category, expires FROM news ORDER BY created DESC LIMIT ?`

	rows, err := m.DB.Query(stmt, limit)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var newsList []News

	for rows.Next() {
		var news News
		err := rows.Scan(&news.ID, &news.Title, &news.Content, &news.Author, &news.Created, &news.Category, &news.Expires)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("error scanning rows: %w", err)
		}

		newsList = append(newsList, news)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return newsList, nil
}

func (m *NewsModel) ByCategory(category string) ([]News, error) {
	stmt := `SELECT id, title, content, author, created, category, expires FROM news WHERE category = ? ORDER BY created DESC`

	rows, err := m.DB.Query(stmt, category)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var newsList []News

	for rows.Next() {
		var news News
		err := rows.Scan(&news.ID, &news.Title, &news.Content, &news.Author, &news.Created, &news.Category, &news.Expires)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("error scanning rows: %w", err)
		}

		newsList = append(newsList, news)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return newsList, nil
}
