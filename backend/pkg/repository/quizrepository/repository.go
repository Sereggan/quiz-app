package quizrepository

import (
	"github.com/Sereggan/quiz-app/pkg/repository"
	"github.com/jackc/pgx/v4"
)

type Repository struct {
	conn           *pgx.Conn
	quizRepository *QuizRepository
}

func New(conn *pgx.Conn) *Repository {
	return &Repository{
		conn: conn,
	}
}

func (r *Repository) Quiz() repository.QuizRepository {
	if r.quizRepository != nil {
		return r.quizRepository
	}

	r.quizRepository = &QuizRepository{
		repository: r,
	}

	return r.quizRepository
}
