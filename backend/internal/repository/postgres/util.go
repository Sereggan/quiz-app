package postgres

import (
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
		err := rows.Scan(&id, &description, &answer)
		if err != nil {
			return nil, fmt.Errorf("could not parse quzzes, error: %s", err)
		}

		quizzes = append(quizzes, &model.Quiz{
			Id:          id,
			Description: description,
			Answer:      answer,
		})
	}

	return quizzes, nil
}
