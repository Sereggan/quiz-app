package main

import (
	"github.com/Sereggan/quiz-app/internal/server/restserver"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	s := restserver.New()

	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
