package main

import (
	"github.com/Sereggan/quiz-app/pkg/config"
	"github.com/Sereggan/quiz-app/pkg/server/restserver"
	"log"
)

var (
	configPath = "./config/config.yaml"
)

func main() {
	appConfig, err := config.New(configPath)
	if err != nil {
		log.Fatal(err)
	}

	s := restserver.New(appConfig)
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
