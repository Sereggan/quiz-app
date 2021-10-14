package service

import (
	"fmt"
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/Sereggan/quiz-app/internal/repository"
	"strings"
)

type QuizService struct {
	quizRepository repository.QuizRepository
}

func NewQuizService(repo repository.Quiz) *QuizService {
	return &QuizService{
		quizRepository: repo,
	}
}

func (q *QuizService) Create(quiz *model.Quiz) error {
	err := q.Create(quiz)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("New quiz was creater: %+v\n", quiz)
	return nil
}

func (q *QuizService) GetById(id int) (*model.Quiz, error) {
	quiz, err := q.quizRepository.Find(id)
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

func (q *QuizService) Delete(id int) error {
	err := q.Delete(id)
	if err != nil {
		fmt.Println(&err)
		return err
	}
	fmt.Printf("Quiz was deleted, quiz id: %d\n", id)

	return nil
}

func (q *QuizService) SolveQuiz(solution *model.Solution) (*model.SolutionResponse, error) {
	quiz, err := q.quizRepository.Find(solution.QuizId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if isRightAnswer(quiz.Answer, solution.Solution) {
		fmt.Printf("Quiz was successfullty solved, quiz id: %d\n", solution.QuizId)
		return &model.SolutionResponse{IsRight: true}, nil
	}
	fmt.Printf("Quiz was not solved, quiz id: %d\n", solution.QuizId)

	return &model.SolutionResponse{IsRight: false}, nil
}

func isRightAnswer(answer string, solution string) bool {
	return strings.ToLower(answer) == strings.ToLower(solution)
}
