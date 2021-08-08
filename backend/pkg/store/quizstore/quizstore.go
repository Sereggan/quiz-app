package quizstore

import (
	"database/sql"
	"log"
)

type QuizStore struct {
	db *sql.DB
}

func New(address string) *QuizStore {
	db, err := newDB(address)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return &QuizStore{
		db: db,
	}
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
