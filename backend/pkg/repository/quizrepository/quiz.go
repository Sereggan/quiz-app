package quizrepository

type Quiz struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Answer      string `json:"answer"`
}
