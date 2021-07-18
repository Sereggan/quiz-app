package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router   *mux.Router
	logger   *logrus.Logger
	basePath string
}

func NewServer(config *Config) *server {
	serverAddr := config.Server.Host + ":" + config.Server.Port
	s := &server{
		basePath: serverAddr,
		router:   mux.NewRouter(),
		logger:   logrus.New(),
	}

	setLoggerLevel(s.logger, config.LogLevel)
	s.configureRouter()

	return s
}

func (s *server) Start() error {
	return http.ListenAndServe(s.basePath, s)
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

func setLoggerLevel(logger *logrus.Logger, level string) {
	ll, err := logrus.ParseLevel(level)
	if err != nil {
		ll = logrus.DebugLevel
	}

	logrus.SetLevel(ll)
}
