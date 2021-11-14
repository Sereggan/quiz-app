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
	router.Use(CORSMiddleware())
	router.Use(jsonContentTypeMiddleware)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
		auth.POST("/logout", h.logOut)
	}

	api := router.Group("/api", h.userIdentity)
	{
		quiz := api.Group("/quiz")
		{
			quiz.POST("/", h.CreateQuiz)
			quiz.GET("/", h.GetAllQuizzes)
			quiz.GET("/:id", h.GetQuiz)
			quiz.DELETE("/:id", h.DeleteQuiz)
			quiz.POST("/solve", h.SolveQuiz)
			quiz.PUT("/", h.UpdateQuiz)
		}
	}

	return router
}
