package postgres

import (
	"context"
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/jackc/pgx/v4"
)

type UserRepository struct {
	conn *pgx.Conn
}

func (a *UserRepository) Create(context context.Context, user *model.User) (int, error) {
	var id int
	err := a.conn.QueryRow(context,
		"INSERT INTO users (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id",
		user.Name,
		user.Username,
		user.Password).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *UserRepository) Find(context context.Context, username string, password string) (model.User, error) {
	user := model.User{
		Username: username,
		Password: password,
	}

	err := a.conn.QueryRow(context,
		"SELECT id, name from users where username=$1 AND password_hash=$2", username, password).
		Scan(&user.Id,
			&user.Name)

	return user, err
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{conn: conn}
}
