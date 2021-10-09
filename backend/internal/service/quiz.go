package service

import (
	"fmt"
	"github.com/Sereggan/quiz-app/internal/config"
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/Sereggan/quiz-app/internal/repository"
	"github.com/Sereggan/quiz-app/internal/repository/postgres"
	"log"
	"strings"
)

type QuizService struct {
	repository repository.Repository
}

func New() QuizService {
	configMap := config.New()
	conn, err := repository.GetConnection(configMap.DbAddress)

	if err != nil {
		log.Fatalf("Could not create connection to: %s", configMap.DbAddress)
	}

	return QuizService{
		repository: postgres.New(conn),
	}
}

func (q *QuizService) Create(quiz *model.Quiz) error {
	err := q.repository.Quiz().Create(quiz)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("New quiz was creater: %+v\n", quiz)
	return nil
}

func (q *QuizService) GetById(id int) (*model.Quiz, error) {
	quiz, err := q.repository.Quiz().Find(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("Quiz was found: %+v\n", *quiz)

	return quiz, nil
}

func (q *QuizService) GetAll() ([]*model.Quiz, error) {

	quizzes, err := q.repository.Quiz().FindAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("%d quizzes were found\n", len(quizzes))

	return quizzes, nil
}

func (q *QuizService) Delete(id int) error {
	err := q.repository.Quiz().Delete(id)
	if err != nil {
		fmt.Println(&err)
		return err
	}
	fmt.Printf("Quiz was deleted, quiz id: %d\n", id)

	return nil
}

func (q *QuizService) SolveQuiz(solution *Solution) (*SolutionResponse, error) {
	quiz, err := q.repository.Quiz().Find(solution.QuizId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if isRightAnswer(quiz.Answer, solution.Solution) {
		fmt.Printf("Quiz was successfullty solved, quiz id: %d\n", solution.QuizId)
		return &SolutionResponse{isRight: true}, nil
	}
	fmt.Printf("Quiz was not solved, quiz id: %d\n", solution.QuizId)

	return &SolutionResponse{isRight: false}, nil
}

func isRightAnswer(answer string, solution string) bool {
	return strings.ToLower(answer) == strings.ToLower(solution)
}
