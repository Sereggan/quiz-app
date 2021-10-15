package repository

import (
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/Sereggan/quiz-app/internal/repository/postgres"
	"github.com/jackc/pgx/v4"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Quiz interface {
	Create(quiz *model.Quiz) error
	Find(id int) (*model.Quiz, error)
	FindAll() ([]*model.Quiz, error)
	Delete(id int) error
}

type Repository struct {
	Quiz
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{
		Quiz: postgres.NewQuizRepository(conn),
	}
}
