package server

import (
	config2 "github.com/Sereggan/quiz-app/pkg/config"
)

type server interface {
	Start() error
	NewServer(config *config2.Config) *server
}
