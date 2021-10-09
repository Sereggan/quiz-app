package handler

import (
	"encoding/json"
	"github.com/Sereggan/quiz-app/internal/server/restserver"
	"github.com/Sereggan/quiz-app/internal/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type QuizHandler struct {
	quizService service.QuizService
}

func New() QuizHandler {
	return QuizHandler{
		quizService: service.New(),
	}
}

func (s *QuizHandler) HandleCreate() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		newQuiz, err := restserver.RetrieveQuiz(request)
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

		writer.WriteHeader(http.StatusNoContent)
	}
}

func (s *QuizHandler) HandleSolve() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		solution, err := restserver.ExtractSolution(request)
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
