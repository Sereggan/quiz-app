package handler

import (
	"encoding/json"
	"github.com/Sereggan/quiz-app/pkg/server/util"
	"github.com/Sereggan/quiz-app/pkg/service/quizservice"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type QuizHandler struct {
	quizService *quizservice.QuizService
}

func New() *QuizHandler {
	return &QuizHandler{
		quizService: quizservice.New(),
	}
}

func (s *QuizHandler) HandlePost() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		newQuiz, err := util.ExtractQuiz(request)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		createdQuiz, err := s.quizService.Create(newQuiz)

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

func (s *QuizHandler) HandleGetQuizById() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		quiz, err := s.quizService.Get(id)

		if err != nil {
			writer.WriteHeader(400)
			return
		}

		err = json.NewEncoder(writer).Encode(&quiz)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		writer.WriteHeader(http.StatusCreated)
	}
}

func (s *QuizHandler) HandleGet() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		quiz, err := s.quizService.GetAll()

		if err != nil {
			writer.WriteHeader(400)
			return
		}

		err = json.NewEncoder(writer).Encode(&quiz)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		writer.WriteHeader(http.StatusCreated)
	}
}
