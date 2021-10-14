package service

import (
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/Sereggan/quiz-app/internal/repository"
)

type Quiz interface {
	Create(quiz *model.Quiz) error
	GetById(id int) (*model.Quiz, error)
	GetAll() ([]*model.Quiz, error)
	Delete(id int) error
	SolveQuiz(solution *model.Solution) (*model.SolutionResponse, error)
}

type Service struct {
	Quiz
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Quiz: NewQuizService(repos.Quiz),
	}
}
