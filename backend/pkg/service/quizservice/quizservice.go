package quizservice

import (
	"fmt"
	"github.com/Sereggan/quiz-app/pkg/config"
	"github.com/Sereggan/quiz-app/pkg/repository"
	"github.com/Sereggan/quiz-app/pkg/repository/quizrepository"
	"log"
	"strings"
)

type QuizService struct {
	quizRepository repository.Repository
}

func New() QuizService {
	configMap := config.New()
	conn, err := repository.GetConnection(configMap.DbAddress)

	if err != nil {
		log.Fatalf("Could not create connection to: %s", configMap.DbAddress)
	}

	return QuizService{
		quizRepository: quizrepository.New(conn),
	}
}

func (q *QuizService) Create(quiz *quizrepository.Quiz) error {
	quiz, err := q.quizRepository.Quiz().Create(quiz)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("New quiz was creater: %+v\n", *quiz)

	return nil
}

func (q *QuizService) GetById(id int) (*quizrepository.Quiz, error) {
	quiz, err := q.quizRepository.Quiz().Find(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("Quiz was found: %+v\n", *quiz)

	return quiz, nil
}

func (q *QuizService) GetAll() ([]*quizrepository.Quiz, error) {

	quizzes, err := q.quizRepository.Quiz().FindAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("%d quizzes were found\n", len(quizzes))

	return quizzes, nil
}

func (q *QuizService) Delete(id int) error {
	err := q.quizRepository.Quiz().Delete(id)
	if err != nil {
		fmt.Println(&err)
		return err
	}
	fmt.Printf("Quiz was deleted, quiz id: %d\n", id)

	return nil
}

func (q *QuizService) SolveQuiz(solution *Solution) (*SolutionResponse, error) {
	quiz, err := q.quizRepository.Quiz().Find(solution.QuizId)
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
