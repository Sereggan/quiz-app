package quizrepository

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestQuizRepository_SaveQuiz(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()

}
