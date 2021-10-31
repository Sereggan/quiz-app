package service

import (
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/Sereggan/quiz-app/internal/repository"
)

type Quiz interface {
	Create(*model.Quiz) error
	GetById(int) (*model.Quiz, error)
	GetAll() ([]*model.Quiz, error)
	Delete(int, int) error
	Update(*model.Quiz) error
	SolveQuiz(*model.Solution) (*model.SolutionResponse, error)
}

type Auth interface {
	CreateUser(*model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	LogOut(int) error
	ParseToken(string) (int, error)
}

type Service struct {
	Quiz
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Quiz: NewQuizService(repos.Quiz),
		Auth: NewAuthService(repos.User, repos.TokenClient),
	}
}
