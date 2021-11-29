package service

import (
	"context"
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/Sereggan/quiz-app/internal/repository"
	"github.com/sirupsen/logrus"
	"strings"
)

type QuizService struct {
	repository repository.Quiz
}

func NewQuizService(repo repository.Quiz) *QuizService {
	return &QuizService{
		repository: repo,
	}
}

func (q *QuizService) Create(context context.Context, quiz *model.Quiz) error {
	err := q.repository.Create(quiz)
	if err != nil {
		logrus.Errorf("Failed to create quiz: %v, err: %s", *quiz, err.Error())
		return err
	}
	logrus.Printf("New quiz was creater: %v\n", *quiz)
	return nil
}

func (q *QuizService) GetById(context context.Context, id int) (*model.Quiz, error) {
	quiz, err := q.repository.Find(id)
	if err != nil {
		logrus.Errorf("Failed to get quiz by id: %d, err: %s", id, err.Error())
		return nil, err
	}
	logrus.Printf("Quiz was found: %+v\n", *quiz)

	return quiz, nil
}

func (q *QuizService) GetAll(context context.Context) ([]*model.Quiz, error) {

	quizzes, err := q.repository.FindAll()
	if err != nil {
		logrus.Errorf("Failed to get all quizzes, err: %s", err.Error())
		return nil, err
	}
	logrus.Printf("%d quizzes were found\n", len(quizzes))

	return quizzes, nil
}

func (q *QuizService) Update(context context.Context, quiz *model.Quiz) error {
	err := q.repository.Update(quiz)
	if err != nil {
		logrus.Errorf("Failed to update quiz: %v, err: %s", *quiz, err.Error())
		return err
	}
	logrus.Printf("Quiz was updated: %v\n", *quiz)
	return nil
}

func (q *QuizService) Delete(context context.Context, quizId int, userId int) error {
	err := q.repository.Delete(quizId, userId)
	if err != nil {
		logrus.Errorf("Failed to delete quiz with id: %d, err: %s", quizId, err.Error())
		return err
	}
	logrus.Printf("Quiz was deleted, quiz id: %d, user_id: %d\n", quizId, userId)

	return nil
}

func (q *QuizService) SolveQuiz(context context.Context, solution *model.Solution) (*model.SolutionResponse, error) {
	quiz, err := q.repository.Find(solution.QuizId)
	if err != nil {
		logrus.Errorf("Failed to find quiz with id: %d, err: %s", solution.QuizId, err.Error())
		return nil, err
	}

	if isRightAnswer(quiz.Answer, solution.Solution) {
		logrus.Printf("Quiz was successfullty solved, quiz id: %d\n", solution.QuizId)
		return &model.SolutionResponse{IsRight: true}, nil
	}
	logrus.Printf("Quiz was not solved, quiz id: %d\n", solution.QuizId)

	return &model.SolutionResponse{IsRight: false}, nil
}

func isRightAnswer(answer string, solution string) bool {
	return strings.ToLower(answer) == strings.ToLower(solution)
}
