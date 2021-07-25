package restserver

import (
	"fmt"
	config2 "github.com/Sereggan/quiz-app/pkg/config"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	Logger *logrus.Logger
)

type server struct {
	router   *mux.Router
	logger   *logrus.Logger
	basePath string
}

func New(config *config2.Config) *server {

	s := &server{
		basePath: config.Server.Address,
		router:   mux.NewRouter(),
		logger:   logrus.New(),
	}

	setLoggerLevel(s.logger, config.LogLevel)

	s.configureRouter()

	s.logger.Debugln("Server configured successfully")
	Logger = s.logger
	return s
}

func (s *server) Start() error {
	s.logger.Debugf("Starting server on: %s, with debug level: %s", s.basePath, s.logger.Level)
	fmt.Println(s.logger.GetLevel())
	return http.ListenAndServe(s.basePath, s)
}

func (s *server) configureRouter() {
	s.logger.Debugf("Server starts configuring routes")

	s.router.Use(setJsonContentTypeMiddleware)
	s.router.Use(mux.CORSMethodMiddleware(s.router))
	// For develop process, will be deleted in future
	s.router.Use(LoggingMiddleware)

	s.router.HandleFunc("/hello", s.HandleUsersCreate()).Methods(http.MethodGet)
}

func (s *server) HandleUsersCreate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello"))
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func setLoggerLevel(logger *logrus.Logger, level string) {
	ll, err := logrus.ParseLevel(level)
	if err != nil {
		logger.Warningln(err)
		ll = logrus.DebugLevel
	}

	logger.SetLevel(ll)
}

func setJsonContentTypeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
