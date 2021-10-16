package handler

import (
	"github.com/Sereggan/quiz-app/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{
		services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		quiz := api.Group("/quiz")
		{
			quiz.POST("/", h.CreateQuiz)
			quiz.GET("/", h.GetAllQuizzes)
			quiz.GET("/:id", h.GetQuiz)
			quiz.DELETE("/:id", h.DeleteQuiz)
			quiz.POST("/solve", h.SolveQuiz)
		}
	}

	return router
}
