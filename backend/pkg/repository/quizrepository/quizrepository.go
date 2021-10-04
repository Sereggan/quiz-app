package quizrepository

import (
	"context"
)

type QuizRepository struct {
	repository *Repository
}

func (r *QuizRepository) Create(quiz *Quiz) error {

	err := r.repository.conn.QueryRow(context.Background(),
		"INSERT INTO quiz (description, answer) VALUES ($1, $2) RETURNING id",
		quiz.Description,
		quiz.Answer).Scan(&quiz.Id)

	if err != nil {
		return err
	}
	return nil
}

func (r *QuizRepository) Find(id int) (*Quiz, error) {

	quiz := &Quiz{}
	err := r.repository.conn.QueryRow(context.Background(),
		"SELECT id, description, answer from quiz where id=$1", id).
		Scan(&quiz.Id,
			&quiz.Description,
			&quiz.Answer)

	if err != nil {
		return nil, err
	}

	return quiz, nil
}

func (r *QuizRepository) FindAll() ([]*Quiz, error) {

	rows, err := r.repository.conn.Query(context.Background(),
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

func (r *QuizRepository) Delete(id int) error {

	commandTag, err := r.repository.conn.Exec(context.Background(), "DELETE from quiz where id=$1", id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return &RepositoryError{message: "No quizzes to delete"}
	}

	return nil
}
