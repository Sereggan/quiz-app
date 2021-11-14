package model

type Quiz struct {
	Id          int    `json:"id"`
	Description string `json:"description" binding:"required"`
	Answer      string `json:"answer" binding:"required"`
	UserId      int    `json:"userId"`
}
