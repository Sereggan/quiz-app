package repository

import (
	"context"
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/Sereggan/quiz-app/internal/repository/postgres"
	tokenrepository "github.com/Sereggan/quiz-app/internal/repository/redis"
	"github.com/go-redis/redis/v8"
	"time"

	"github.com/jackc/pgx/v4"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Quiz interface {
	Create(context context.Context, quiz *model.Quiz) error
	Find(context context.Context, id int) (*model.Quiz, error)
	FindAll(context context.Context) ([]*model.Quiz, error)
	Update(context context.Context, quiz *model.Quiz) error
	Delete(context context.Context, quizId int, userId int) error
}

type User interface {
	Create(context context.Context, user *model.User) (int, error)
	Find(context context.Context, username string, password string) (model.User, error)
}

type TokenClient interface {
	Set(context context.Context, key string, value interface{}, ttl time.Duration) error
	Get(context context.Context, key string) (value interface{}, err error)
	Delete(context context.Context, key string) error
}

type Repository struct {
	Quiz
	User
	TokenClient
}

func NewRepository(connPostgres *pgx.Conn, redisClient *redis.Client) *Repository {
	return &Repository{
		postgres.NewQuizRepository(connPostgres),
		postgres.NewUserRepository(connPostgres),
		tokenrepository.NewTokenCLient(redisClient),
	}
}
