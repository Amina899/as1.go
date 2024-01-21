package mysql

import (
	"aitu/aitunews/pkg/models"
	"database/sql"
	"errors"
	"time"
)

type NewsModel struct {
	DB *sql.DB
}

func (m *NewsModel) InsertNews(title, content, author string, created time.Time, category string) (int, error) {
	stmt := `INSERT INTO thenews (title, content, author, created, category) VALUES (?, ?, ?, ?, ?)`
	result, err := m.DB.Exec(stmt, title, content, author, created, category)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
func (m *NewsModel) Get(id int) (*models.News, error) {
	stmt := `SELECT id, title, content, author, created, category FROM thenews WHERE id =?`
	row := m.DB.QueryRow(stmt, id)
	s := &models.News{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Author, &s.Created, &s.Category)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// Latest This will return the 10 most recently created snippets.
func (m *NewsModel) Latest() ([]*models.News, error) {
	return nil, nil
}
