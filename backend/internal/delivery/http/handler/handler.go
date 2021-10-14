package handler

import "github.com/Sereggan/quiz-app/internal/service"

type Handler struct {
	service *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{
		service: services,
	}
}
