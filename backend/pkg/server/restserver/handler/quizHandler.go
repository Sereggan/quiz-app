package handler

import (
	"encoding/json"
	"github.com/Sereggan/quiz-app/pkg/server/util"
	"github.com/Sereggan/quiz-app/pkg/service/quizservice"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"strconv"
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

func (s *QuizHandler) HandleGetQuiz() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		quiz, err := s.quizService.GetQuiz(id)

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
		quiz, err := s.quizService.GetAllQuizzes()

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

func defineLimit(values *url.Values) int {
	limit := values.Get("limit")
	var limitInt int
	if limit == "" {
		limitInt = 5
	} else {
		limitInt, _ = strconv.Atoi(limit)
	}
	return limitInt
}
