package main

import (
	"context"
	"github.com/Sereggan/quiz-app/internal/config"
	"github.com/Sereggan/quiz-app/internal/delivery/http/handler"
	"github.com/Sereggan/quiz-app/internal/repository"
	"github.com/Sereggan/quiz-app/internal/repository/postgres"
	"github.com/Sereggan/quiz-app/internal/repository/redis"
	"github.com/Sereggan/quiz-app/internal/server/restserver"
	"github.com/Sereggan/quiz-app/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		logrus.Print("No .env file found")
	}
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg := config.New()

	conn, err := postgres.GetConnection(cfg.DbAddress)
	if err != nil {
		logrus.Fatal(err)
	}

	client, err := redis.GetClient(cfg.RedisAddress)
	if err != nil {
		logrus.Fatal(err)
	}

	repos := repository.NewRepository(conn, client)
	services := service.NewService(repos)
	handlers := handler.New(services)
	srv := new(restserver.Server)

	go func() {
		err = srv.Run(handlers.InitRoutes())
		if err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Printf("QuizApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("QuizApp Shutting Down")

	// GraceFul shutdown
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := conn.Close(context.Background()); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
