package restserver

import (
	"encoding/json"
	"fmt"
	"github.com/Sereggan/quiz-app/pkg/config"
	"github.com/Sereggan/quiz-app/pkg/server/util"
	"github.com/Sereggan/quiz-app/pkg/service/quizservice"
	"github.com/gorilla/mux"
	"net/http"
)

type server struct {
	basePath    string
	router      *mux.Router
	quizService *quizservice.QuizService
}

func New(config *config.Config) *server {

	s := &server{
		basePath:    config.Server.Address,
		router:      mux.NewRouter(),
		quizService: quizservice.New(config.DB.Address),
	}

	s.configureRouter()

	fmt.Println("Server configured successfully")
	return s
}

func (s *server) Start() error {
	fmt.Printf("Starting server on: %s\n", s.basePath)
	return http.ListenAndServe(s.basePath, s)
}

func (s *server) configureRouter() {
	fmt.Println("Server starts configuring routes")

	s.router.Use(setJsonContentTypeMiddleware)
	s.router.Use(mux.CORSMethodMiddleware(s.router))

	s.router.HandleFunc("/quiz", addQuiz(s.quizService)).Methods(http.MethodPost)
}

func addQuiz(quizService *quizservice.QuizService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		newQuiz, err := util.ExtractQuiz(request)
		if err != nil {
			writer.WriteHeader(400)
			return
		}
		createdQuiz, err := quizService.CreateQuiz(newQuiz)

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

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func setJsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
