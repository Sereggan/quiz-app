package handler

import (
	"encoding/json"
	"github.com/Sereggan/quiz-app/pkg/server/util"
	"github.com/Sereggan/quiz-app/pkg/service/quizservice"
	"net/http"
)

type QuizHandler struct {
	quizService *quizservice.QuizService
}

func New(databaseURL string) *QuizHandler {
	return &QuizHandler{
		quizService: quizservice.New(databaseURL),
	}
}

func (s *QuizHandler) HandlePost() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		newQuiz, err := util.ExtractQuiz(request)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		createdQuiz, err := s.quizService.CreateQuiz(newQuiz)

		if err != nil {
			writer.WriteHeader(400)
			return
		}

		err = json.NewEncoder(writer).Encode(&createdQuiz)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		writer.WriteHeader(http.StatusCreated)
	}
}
