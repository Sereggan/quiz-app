package quizservice

import (
	"fmt"
	"github.com/Sereggan/quiz-app/pkg/model"
	"github.com/Sereggan/quiz-app/pkg/repository/quizrepository"
)

type QuizService struct {
	quizRepository *quizrepository.QuizRepository
}

func New() *QuizService {
	repository := quizrepository.New()

	return &QuizService{
		quizRepository: repository,
	}
}

func (q *QuizService) Create(quiz *model.Quiz) (*model.Quiz, error) {
	savedQuiz, err := q.quizRepository.Add(quiz)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("New quiz was creater: %+v\n", *savedQuiz)

	return savedQuiz, nil
}

func (q *QuizService) Get(id int) (*model.Quiz, error) {
	quiz, err := q.quizRepository.FindById(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("Quiz was found: %+v\n", *quiz)

	return quiz, nil
}

func (q *QuizService) GetAll() ([]*model.Quiz, error) {

	quizzes, err := q.quizRepository.FindAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("%d quizzes were found\n", len(quizzes))

	return quizzes, nil
}
