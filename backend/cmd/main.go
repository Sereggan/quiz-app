package main

import (
	config3 "github.com/Sereggan/quiz-app/pkg/config"
	"github.com/Sereggan/quiz-app/pkg/server/restserver"
	"log"
)

var (
	configPath = "./config/config.yaml"
)

func main() {
	config, err := config3.New(configPath)
	if err != nil {
		log.Fatal(err)
	}

	s := restserver.New(config)
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
