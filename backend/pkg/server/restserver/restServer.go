package restserver

import (
	"fmt"
	"github.com/Sereggan/quiz-app/pkg/config"
	"github.com/Sereggan/quiz-app/pkg/server/restserver/handler"
	"github.com/gorilla/mux"
	"net/http"
)

type server struct {
	basePath    string
	router      *mux.Router
	quizHandler *handler.QuizHandler
}

func New(config *config.Config) *server {

	s := &server{
		basePath:    config.Server.Address,
		router:      mux.NewRouter(),
		quizHandler: handler.New(config.DB.Address),
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

	s.router.HandleFunc("/quiz", s.quizHandler.HandlePost()).Methods(http.MethodPost)
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
