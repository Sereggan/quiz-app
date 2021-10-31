package postgres

import (
	"context"
	"fmt"
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type QuizRepository struct {
	conn *pgx.Conn
}

func NewQuizRepository(conn *pgx.Conn) *QuizRepository {
	return &QuizRepository{conn: conn}
}

func (r *QuizRepository) Create(quiz *model.Quiz) error {

	err := r.conn.QueryRow(context.Background(),
		"INSERT INTO quiz (description, answer, user_id) VALUES ($1, $2, $3) RETURNING id",
		quiz.Description,
		quiz.Answer,
		quiz.UserId).Scan(&quiz.Id)

	if err != nil {
		return err
	}
	return nil
}

func (r *QuizRepository) Find(id int) (*model.Quiz, error) {

	quiz := &model.Quiz{}
	err := r.conn.QueryRow(context.Background(),
		"SELECT id, description, answer, user_id from quiz where id=$1", id).
		Scan(&quiz.Id,
			&quiz.Description,
			&quiz.Answer,
			&quiz.UserId)

	if err != nil {
		return nil, fmt.Errorf("could not find value in databse, error: %s, quiz.Id: %d", err, quiz.Id)
	}

	return quiz, nil
}

func (r *QuizRepository) FindAll() ([]*model.Quiz, error) {

	rows, err := r.conn.Query(context.Background(),
		"SELECT id, description, answer, user_id FROM quiz")
	if err != nil {
		return nil, fmt.Errorf("could not find all quizzes in databse, error: %s", err)
	}

	quizzes, err := getQuizzesAsSlice(rows)
	if err != nil {
		return nil, err
	}

	return quizzes, nil
}

func (r *QuizRepository) Update(quiz *model.Quiz) error {
	_, err := r.conn.Exec(context.Background(),
		"UPDATE quiz SET description = $1,"+
			" answer = $2 WHERE id = $3",
		quiz.Description,
		quiz.Answer,
		quiz.Id)

	if err != nil {
		return err
	}
	return nil
}

func (r *QuizRepository) Delete(quizId int, userId int) error {

	commandTag, err := r.conn.Exec(context.Background(), "DELETE from quiz where id=$1 and user_id = $2", quizId, userId)
	if err != nil {
		return fmt.Errorf("could not delete quiz, error: %s, quiz.Id: %d, user_id: %d", err, quizId, userId)
	}
	if commandTag.RowsAffected() != 1 {
		return errors.Wrap(err, "No quizzes to delete")
	}

	return nil
}
