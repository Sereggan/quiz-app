package quizservice

import (
	"fmt"
	"github.com/Sereggan/quiz-app/pkg/repository/quizrepository"
	"strings"
)

type QuizService struct {
	quizRepository *quizrepository.QuizRepository
}

func New() *QuizService {
	return &QuizService{
		quizRepository: quizrepository.New(),
	}
}

func (q *QuizService) Create(quiz *quizrepository.Quiz) error {
	quiz, err := q.quizRepository.Add(quiz)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("New quiz was creater: %+v\n", *quiz)

	return nil
}

func (q *QuizService) GetById(id int) (*quizrepository.Quiz, error) {
	quiz, err := q.quizRepository.FindById(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("Quiz was found: %+v\n", *quiz)

	return quiz, nil
}

func (q *QuizService) GetAll() ([]*quizrepository.Quiz, error) {

	quizzes, err := q.quizRepository.FindAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("%d quizzes were found\n", len(quizzes))

	return quizzes, nil
}

func (q *QuizService) Delete(id int) error {
	err := q.quizRepository.DeleteById(id)
	if err != nil {
		fmt.Println(&err)
		return err
	}
	fmt.Printf("Quiz was deleted, quiz id: %d\n", id)

	return nil
}

func (q *QuizService) SolveQuiz(solution *Solution) (*SolutionResponse, error) {
	quiz, err := q.quizRepository.FindById(solution.QuizId)
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
