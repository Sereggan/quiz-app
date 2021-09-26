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

func (s *QuizHandler) HandleAdd() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		newQuiz, err := util.ExtractQuiz(request)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		err = s.quizService.Create(newQuiz)

		if err != nil {
			writer.WriteHeader(400)
			return
		}

		writer.WriteHeader(http.StatusCreated)
	}
}

func (s *QuizHandler) HandleGetById() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		quiz, err := s.quizService.GetById(id)

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

func (s *QuizHandler) HandleGetAll() func(http.ResponseWriter, *http.Request) {
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

func (s *QuizHandler) HandleDeleteById() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		err = s.quizService.Delete(id)

		if err != nil {
			writer.WriteHeader(400)
			return
		}

		if err != nil {
			writer.WriteHeader(400)
			return
		}

		writer.WriteHeader(http.StatusNoContent)
	}
}

func (s *QuizHandler) HandleSolve() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		solution, err := util.ExtractSolution(request)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		result, err := s.quizService.SolveQuiz(solution)

		if err != nil {
			writer.WriteHeader(400)
			return
		}

		err = json.NewEncoder(writer).Encode(&result)

		if err != nil {
			writer.WriteHeader(400)
			return
		}

		writer.WriteHeader(http.StatusAccepted)
	}
}
