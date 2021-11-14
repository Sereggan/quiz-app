build:
	go build -o bin/server cmd/quizapp/main.go

run:
	go run cmd/quizapp/main.go

test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
