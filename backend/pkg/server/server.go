package server

import (
	"github.com/Sereggan/quiz-app/pkg/config"
)

type server interface {
	Start() error
	NewServer(config *config.Config) *server
}
