package questions

import (
	"github.com/DTS-STN/question-priority-service/models"
)

// QuestionInterface for getting questions and loading questions from file
type QuestionInterface interface {
	Questions() []models.Question
	loadQuestions() ([]models.Question, error)
}

// QuestionServiceStruct is a struct used for passing in the questions file
type QuestionServiceStruct struct {
	Filename string
}

// QuestionService is an instance of QuestionInterface
var QuestionService QuestionInterface
