package quizrepository

import (
	"github.com/jackc/pgx/v4"
)

func getQuizzesAsSlice(rows pgx.Rows) ([]*Quiz, error) {
	var quizzes []*Quiz

	for rows.Next() {
		var id int
		var description string
		var answer string
		err := rows.Scan(&id, &description, &answer)
		if err != nil {
			// handle this error
			return nil, err
		}
		quizzes = append(quizzes, &Quiz{
			Id:          id,
			Description: description,
			Answer:      answer,
		})
	}

	return quizzes, nil
}
