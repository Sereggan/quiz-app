package quizservice

type Solution struct {
	QuizId   int    `json:"quizId"`
	Solution string `json:"solution"`
}
