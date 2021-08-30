package quizrepository

import (
	"context"
	"github.com/Sereggan/quiz-app/pkg/model"
	"github.com/jackc/pgx/v4"
)

func getQuizzesAsSlice(rows pgx.Rows) ([]*model.Quiz, error) {
	var quizzes []*model.Quiz

	for rows.Next() {
		var id int
		var description string
		var answer string
		err := rows.Scan(&id, &description, &answer)
		if err != nil {
			// handle this error
			return nil, err
		}
		quizzes = append(quizzes, &model.Quiz{
			Id:          id,
			Description: description,
			Answer:      answer,
		})
	}

	return quizzes, nil
}

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
