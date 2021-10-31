package repository

import (
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/Sereggan/quiz-app/internal/repository/postgres"
	tokenrepository "github.com/Sereggan/quiz-app/internal/repository/redis"
	"github.com/go-redis/redis/v8"
	"time"

	"github.com/jackc/pgx/v4"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Quiz interface {
	Create(quiz *model.Quiz) error
	Find(id int) (*model.Quiz, error)
	FindAll() ([]*model.Quiz, error)
	Update(quiz *model.Quiz) error
	Delete(quizId int, userId int) error
}

type User interface {
	Create(user *model.User) (int, error)
	Find(username string, password string) (model.User, error)
}

type TokenClient interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string) (value interface{}, err error)
	Delete(key string) error
}

type Repository struct {
	Quiz
	User
	TokenClient
}

func NewRepository(connPostgres *pgx.Conn, redisClient *redis.Client) *Repository {
	return &Repository{
		Quiz:        postgres.NewQuizRepository(connPostgres),
		User:        postgres.NewUserRepository(connPostgres),
		TokenClient: tokenrepository.NewTokenCLient(redisClient),
	}
}
