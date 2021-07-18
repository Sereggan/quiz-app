package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
}

func NewServer() *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
	}

	s.configureRouter()

	return s
}

func (s *server) Start(config *Config) error {
	serverAddr := config.Server.Host + ":" + config.Server.Port
	return http.ListenAndServe(serverAddr, s)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/hello", s.HandleUsersCreate()).Methods(http.MethodGet)
}

func (s *server) HandleUsersCreate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World!")
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
