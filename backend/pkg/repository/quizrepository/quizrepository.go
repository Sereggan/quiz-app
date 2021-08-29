package quizrepository

import (
	"context"
	"fmt"
	"github.com/Sereggan/quiz-app/pkg/model"
	"github.com/jackc/pgx/v4"
	"log"
)

type QuizRepository struct {
	address string
}

func New(databaseURL string) *QuizRepository {
	_, err := getConnection(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully connected to databese on url: %s\n", databaseURL)

	return &QuizRepository{
		address: databaseURL,
	}
}

func (r *QuizRepository) SaveQuiz(quiz *model.Quiz) (*model.Quiz, error) {
	conn, err := getConnection(r.address)

	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	row := conn.QueryRow(context.Background(),
		"INSERT INTO quiz (description, answer) VALUES ($1, $2) RETURNING id",
		quiz.Description,
		quiz.Answer)

	var id int
	err = row.Scan(&id)

	if err != nil {
		return nil, err
	}
	quiz.Id = id

	return quiz, nil
}

//func (r *QuizRepository) GetQuiz(id int) model.Quiz {
//
//}

func getConnection(databaseURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return conn, nil
}
