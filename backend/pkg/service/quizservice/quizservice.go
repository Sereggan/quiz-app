package quizservice

import (
	"fmt"
	"github.com/Sereggan/quiz-app/pkg/model"
	"github.com/Sereggan/quiz-app/pkg/repository/quizrepository"
)

type QuizService struct {
	quizRepository *quizrepository.QuizRepository
}

func New(databaseURL string) *QuizService {
	repository := quizrepository.New(databaseURL)

	return &QuizService{
		quizRepository: repository,
	}
}

func (q *QuizService) CreateQuiz(quiz *model.Quiz) (*model.Quiz, error) {
	savedQuiz, err := q.quizRepository.SaveQuiz(quiz)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("New quiz was creater: %+v\n", *savedQuiz)

	return savedQuiz, nil
}
