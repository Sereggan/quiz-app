package postgres

import (
	"context"
	"fmt"
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/jackc/pgx/v4"
)

func getQuizzesAsSlice(rows pgx.Rows) ([]*model.Quiz, error) {
	var quizzes []*model.Quiz

	for rows.Next() {
		var id int
		var description string
		var answer string
		var user_id int
		err := rows.Scan(&id, &description, &answer)
		if err != nil {
			return nil, fmt.Errorf("could not parse quzzes, error: %s", err)
		}

		quizzes = append(quizzes, &model.Quiz{
			Id:          id,
			Description: description,
			Answer:      answer,
			UserId:      user_id,
		})
	}

	return quizzes, nil
}

func GetConnection(databaseURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}
	return conn, nil
}
