package restserver

import (
	"encoding/json"
	"fmt"
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/Sereggan/quiz-app/internal/service"
	"net/http"
)

func RetrieveQuiz(request *http.Request) (*model.Quiz, error) {

	var newQuiz model.Quiz
	err := json.NewDecoder(request.Body).Decode(&newQuiz)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &newQuiz, nil
}

func ExtractSolution(request *http.Request) (*service.Solution, error) {

	var solution service.Solution
	err := json.NewDecoder(request.Body).Decode(&solution)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &solution, nil
}
