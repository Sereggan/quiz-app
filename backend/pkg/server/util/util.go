package util

import (
	"encoding/json"
	"fmt"
	"github.com/Sereggan/quiz-app/pkg/model"
	"net/http"
)

func ExtractQuiz(request *http.Request) (*model.Quiz, error) {

	var newQuiz model.Quiz
	err := json.NewDecoder(request.Body).Decode(&newQuiz)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &newQuiz, nil
}
