package service

import (
	"context"
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/Sereggan/quiz-app/internal/repository"
)

type Quiz interface {
	Create(context.Context, *model.Quiz) error
	GetById(context.Context, int) (*model.Quiz, error)
	GetAll(context.Context) ([]*model.Quiz, error)
	Delete(context.Context, int, int) error
	Update(context.Context, *model.Quiz) error
	SolveQuiz(context.Context, *model.Solution) (*model.SolutionResponse, error)
}

type Auth interface {
	CreateUser(context.Context, *model.User) (int, error)
	GenerateToken(context.Context, string, string) (string, error)
	LogOut(context.Context, int) error
	ParseToken(context.Context, string) (int, error)
}

type Service struct {
	Quiz
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewQuizService(repos.Quiz),
		NewAuthService(repos.User, repos.TokenClient),
	}
}
