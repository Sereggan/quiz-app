package repository

import (
	"github.com/Sereggan/quiz-app/internal/model"
)

type QuizRepository interface {
	Create(*model.Quiz) error
	Find(int) (*model.Quiz, error)
	FindAll() ([]*model.Quiz, error)
	Delete(int) error
}
