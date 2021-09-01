package quizrepository

import (
	"context"
	"fmt"
	"github.com/Sereggan/quiz-app/pkg/config"
	"github.com/Sereggan/quiz-app/pkg/model"
	"log"
)

type QuizRepository struct {
	address string
}

func New() *QuizRepository {
	config := config.New()

	conn, err := getConnection(config.DbAddress)
	defer conn.Close(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully connected to databese on url: %s\n", config.DbAddress)

	return &QuizRepository{
		address: config.DbAddress,
	}
}

func (r *QuizRepository) Add(quiz *model.Quiz) (*model.Quiz, error) {
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

func (r *QuizRepository) FindById(id int) (*model.Quiz, error) {
	conn, err := getConnection(r.address)

	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	quiz := &model.Quiz{}
	err = conn.QueryRow(context.Background(),
		"SELECT id, description, answer from quiz where id=$1", id).
		Scan(&quiz.Id,
			&quiz.Description,
			&quiz.Answer)
	if err != nil {
		return nil, err
	}
	return quiz, nil
}

func (r *QuizRepository) FindAll() ([]*model.Quiz, error) {
	conn, err := getConnection(r.address)

	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		"SELECT id, description, answer FROM quiz")
	if err != nil {
		return nil, err
	}

	quizzes, err := getQuizzesAsSlice(rows)
	if err != nil {
		return nil, err
	}
	return quizzes, nil
}
