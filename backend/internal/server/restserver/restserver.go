package restserver

import (
	"context"
	"fmt"
	"github.com/Sereggan/quiz-app/internal/config"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler) error {
	cfg := config.New()

	s.httpServer = &http.Server{
		Addr:         cfg.ServerAddress,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server started successfully")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

//func (s *Server) configureRouter() {
//	fmt.Println("Server starts configuring routes")
//
//	s.router.Use(handler.setJsonContentTypeMiddleware)
//	s.router.Use(mux.CORSMethodMiddleware(s.router))
//
//	s.router.HandleFunc("/quiz", s.handler.CreateQuiz()).Methods(http.MethodPost)
//	s.router.HandleFunc("/quiz/solve", s.handler.SolveQuiz()).Methods(http.MethodPost)
//	s.router.HandleFunc("/quiz", s.handler.GetAllQuizzes()).Methods(http.MethodGet)
//	s.router.HandleFunc("/quiz/{id}", s.handler.GetQuiz()).Methods(http.MethodGet)
//	s.router.HandleFunc("/quiz/{id}", s.handler.DeleteQuiz()).Methods(http.MethodDelete)
//
//}
