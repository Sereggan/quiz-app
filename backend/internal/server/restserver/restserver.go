package restserver

import (
	"context"
	"fmt"
	"github.com/Sereggan/quiz-app/internal/config"
	"github.com/Sereggan/quiz-app/internal/delivery/http/handler"
	"github.com/Sereggan/quiz-app/internal/repository"
	"github.com/Sereggan/quiz-app/internal/service"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
)

type server struct {
	router  *mux.Router
	handler *handler.Handler
}

func New() server {
	cfg := config.New()
	conn, err := getConnection(cfg.DbAddress)
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(conn)
	services := service.NewService(repos)
	handlers := handler.New(services)
	s := server{
		router:  mux.NewRouter(),
		handler: handlers,
	}

	s.configureRouter()
	fmt.Println("Server configured successfully")
	return s
}

func (s *server) Start() error {
	configMap := config.New()

	fmt.Printf("Starting server on: %s\n", configMap.ServerAddress)
	return http.ListenAndServe(configMap.ServerAddress, s)
}

func (s *server) configureRouter() {
	fmt.Println("Server starts configuring routes")

	s.router.Use(setJsonContentTypeMiddleware)
	s.router.Use(mux.CORSMethodMiddleware(s.router))

	s.router.HandleFunc("/quiz", s.handler.CreateQuiz()).Methods(http.MethodPost)
	s.router.HandleFunc("/quiz/solve", s.handler.SolveQuiz()).Methods(http.MethodPost)
	s.router.HandleFunc("/quiz", s.handler.GetAllQuizzes()).Methods(http.MethodGet)
	s.router.HandleFunc("/quiz/{id}", s.handler.GetQuiz()).Methods(http.MethodGet)
	s.router.HandleFunc("/quiz/{id}", s.handler.DeleteQuiz()).Methods(http.MethodDelete)

}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func getConnection(databaseURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return conn, nil
}
