package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Sereggan/quiz-app/internal/model"
	"net/http"
)

func retrieveQuiz(request *http.Request) (*model.Quiz, error) {

	var newQuiz model.Quiz
	err := json.NewDecoder(request.Body).Decode(&newQuiz)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &newQuiz, nil
}

func extractSolution(request *http.Request) (*model.Solution, error) {

	var solution model.Solution
	err := json.NewDecoder(request.Body).Decode(&solution)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &solution, nil
}
