package main

import (
	"github.com/Sereggan/quiz-app/pkg/server"
	"log"
)

var (
	configPath = "./config/config.yaml"
)

func main() {
	config, err := server.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewServer()
	server.Start(config)
}
