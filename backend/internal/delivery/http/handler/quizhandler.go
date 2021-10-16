package handler

import (
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateQuiz(c *gin.Context) {
	var newQuiz model.Quiz
	err := c.BindJSON(&newQuiz)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Quiz.Create(&newQuiz)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok", "Quiz created"})
}

func (h *Handler) GetQuiz(c *gin.Context) {

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	quiz, err := h.services.GetById(id)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, quiz)
}

func (h *Handler) GetAllQuizzes(c *gin.Context) {

	quiz, err := h.services.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, quiz)
}

func (h *Handler) DeleteQuiz(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Delete(id)
	c.JSON(http.StatusOK, statusResponse{"ok", "Quiz deleted"})
}

func (h *Handler) SolveQuiz(c *gin.Context) {
	var solution model.Solution
	err := c.BindJSON(&solution)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := h.services.SolveQuiz(&solution)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, result)

}
