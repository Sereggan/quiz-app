package util

import (
	"encoding/json"
	"fmt"
	"github.com/Sereggan/quiz-app/pkg/repository/quizrepository"
	"github.com/Sereggan/quiz-app/pkg/service/quizservice"
	"net/http"
)

func RetrieveQuiz(request *http.Request) (*quizrepository.Quiz, error) {

	var newQuiz quizrepository.Quiz
	err := json.NewDecoder(request.Body).Decode(&newQuiz)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &newQuiz, nil
}

func ExtractSolution(request *http.Request) (*quizservice.Solution, error) {

	var solution quizservice.Solution
	err := json.NewDecoder(request.Body).Decode(&solution)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &solution, nil
}
