package mysql

import (
	"aitu/aitunews/pkg/models"
	"database/sql"
	"errors"
)

type NewsModel struct {
	DB *sql.DB
}

func (m *NewsModel) Insert(title, content, author, created string) (int, error) {
	stmt := `INSERT INTO news (title, content, author, created) VALUES (?, ?, ?, ?)`
	result, err := m.DB.Exec(stmt, title, content, author)
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
	stmt := `SELECT id, title, content, author, created FROM news WHERE id =?`
	row := m.DB.QueryRow(stmt, id)
	n := &models.News{}
	err := row.Scan(&n.ID, &n.Title, &n.Content, &n.Author, &n.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
		}
	}
	return n, nil
}

// Latest This will return the 10 most recently created NEWS.
func (m *NewsModel) Latest() ([]*models.News, error) {
	stmt := `SELECT id, title, content, author, created FROM news ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	news := []*models.News{}
	for rows.Next() {
		n := &models.News{}
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.Author, &n.Created)
		if err != nil {
			return nil, err
		}
		news = append(news, n)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return news, nil
}
