package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	db *sql.DB
}

func NewMySQLDB(dsn string) (*MySQLDB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &MySQLDB{db: db}, nil
}

func (m *MySQLDB) SaveScore(username string, score int) error {
	_, err := m.db.Exec("INSERT INTO scores (username, score) VALUES (?, ?)", username, score)
	return err
}

func (m *MySQLDB) GetTopScores(limit int) ([]Score, error) {
	rows, err := m.db.Query("SELECT username, score FROM scores ORDER BY score DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []Score
	for rows.Next() {
		var s Score
		if err := rows.Scan(&s.Username, &s.Score); err != nil {
			return nil, err
		}
		scores = append(scores, s)
	}

	return scores, nil
}

func (m *MySQLDB) Close() error {
	return m.db.Close()
}

type Score struct {
	Username string
	Score    int
}
