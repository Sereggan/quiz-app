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

func (q *QuizService) GetQuiz(id int) (*model.Quiz, error) {
	quiz, err := q.quizRepository.GetQuiz(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("Quiz was found: %+v\n", *quiz)

	return quiz, nil
}

func (q *QuizService) GetAllQuizzes() ([]*model.Quiz, error) {

	quizzes, err := q.quizRepository.GetAllQuizzes()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("%d quizzes were found\n", len(quizzes))

	return quizzes, nil
}

func contains(s []int32, str int32) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
