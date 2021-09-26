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

func New() *server {

	configMap := config.New()

	s := &server{
		basePath:    configMap.ServerAddress,
		router:      mux.NewRouter(),
		quizHandler: handler.New(),
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

	s.router.HandleFunc("/quiz", s.quizHandler.HandleAdd()).Methods(http.MethodPost)
	s.router.HandleFunc("/quiz/solve", s.quizHandler.HandleSolve()).Methods(http.MethodPost)
	s.router.HandleFunc("/quiz", s.quizHandler.HandleGetAll()).Methods(http.MethodGet)
	s.router.HandleFunc("/quiz/{id}", s.quizHandler.HandleGetById()).Methods(http.MethodGet)
	s.router.HandleFunc("/quiz/{id}", s.quizHandler.HandleDeleteById()).Methods(http.MethodDelete)

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
