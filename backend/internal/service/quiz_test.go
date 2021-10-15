package service

import (
	"github.com/Sereggan/quiz-app/internal/model"
	"github.com/Sereggan/quiz-app/internal/repository"
	mock_repository "github.com/Sereggan/quiz-app/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	quizId      = 1
	rightAnswer = "Right asnwer"
	description = "Description"
)

func TestQuiz_SolveQuiz(t *testing.T) {

	type mockBehavior func(s *mock_repository.MockQuiz, quiz *model.Quiz)

	testTable := []struct {
		name             string
		quiz             *model.Quiz
		mockBehavior     mockBehavior
		solutionResponse *model.SolutionResponse
		solution         *model.Solution
	}{
		{
			name: "Right answer",
			quiz: &model.Quiz{
				Id:          quizId,
				Description: description,
				Answer:      rightAnswer,
			},
			mockBehavior: func(s *mock_repository.MockQuiz, quiz *model.Quiz) {
				s.EXPECT().Find(quiz.Id).Return(quiz, nil)
			},
			solutionResponse: &model.SolutionResponse{IsRight: true},
			solution: &model.Solution{
				QuizId:   1,
				Solution: rightAnswer,
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			quiz := mock_repository.NewMockQuiz(c)
			testCase.mockBehavior(quiz, testCase.quiz)

			repos := &repository.Repository{Quiz: quiz}
			serv := NewService(repos)

			response, err := serv.SolveQuiz(testCase.solution)

			assert.True(t, err == nil)
			assert.Equal(t, testCase.solutionResponse, response)
		})
	}
}
