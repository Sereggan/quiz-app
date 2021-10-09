package postgres

import (
	"context"
	"fmt"
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/pkg/errors"
)

type QuizRepository struct {
	repository *Repository
}

func (r *QuizRepository) Create(quiz *model.Quiz) error {

	err := r.repository.conn.QueryRow(context.Background(),
		"INSERT INTO quiz (description, answer) VALUES ($1, $2) RETURNING id",
		quiz.Description,
		quiz.Answer).Scan(&quiz.Id)

	if err != nil {
		return err
	}
	return nil
}

func (r *QuizRepository) Find(id int) (*model.Quiz, error) {

	quiz := &model.Quiz{}
	err := r.repository.conn.QueryRow(context.Background(),
		"SELECT id, description, answer from quiz where id=$1", id).
		Scan(&quiz.Id,
			&quiz.Description,
			&quiz.Answer)

	if err != nil {
		return nil, fmt.Errorf("could not find value in databse, error: %s, quiz.Id: %s", err, quiz.Id)
	}

	return quiz, nil
}

func (r *QuizRepository) FindAll() ([]*model.Quiz, error) {

	rows, err := r.repository.conn.Query(context.Background(),
		"SELECT id, description, answer FROM quiz")
	if err != nil {
		return nil, fmt.Errorf("could not find all quizzes in databse, error: %s", err)
	}

	quizzes, err := getQuizzesAsSlice(rows)
	if err != nil {
		return nil, err
	}

	return quizzes, nil
}

func (r *QuizRepository) Delete(id int) error {

	commandTag, err := r.repository.conn.Exec(context.Background(), "DELETE from quiz where id=$1", id)
	if err != nil {
		return fmt.Errorf("could not delete quiz, error: %s, quiz.Id: %s", err, id)
	}
	if commandTag.RowsAffected() != 1 {
		return errors.Wrap(err, "No quizzes to delete")
	}

	return nil
}
