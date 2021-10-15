package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *Handler) CreateQuiz() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		newQuiz, err := retrieveQuiz(request)
		if err != nil {
			writer.WriteHeader(400)
			return
		}

		err = s.Service.Create(newQuiz)
		if err != nil {
			writer.WriteHeader(400)
			return
		}

		writer.WriteHeader(http.StatusCreated)
	}
}

func (s *Handler) GetQuiz() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			writer.WriteHeader(400)
			return
		}

		quiz, err := s.Service.GetById(id)
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

func (s *Handler) GetAllQuizzes() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		quiz, err := s.Service.GetAll()
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

func (s *Handler) DeleteQuiz() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			writer.WriteHeader(400)
			return
		}

		err = s.Service.Delete(id)
		if err != nil {
			writer.WriteHeader(400)
			return
		}

		writer.WriteHeader(http.StatusNoContent)
	}
}

func (s *Handler) SolveQuiz() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		solution, err := extractSolution(request)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		result, err := s.Service.SolveQuiz(solution)

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
