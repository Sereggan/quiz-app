package repository

import "github.com/Sereggan/quiz-app/pkg/repository/quizrepository"

type QuizRepository interface {
	Create(*quizrepository.Quiz) error
	Find(int) (*quizrepository.Quiz, error)
	FindAll() ([]*quizrepository.Quiz, error)
	Delete(int) error
}
