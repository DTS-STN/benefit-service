package questions

import (
	"github.com/DTS-STN/benefit-service/models"
)

// QuestionInterface for getting questions and loading questions from file
type QuestionInterface interface {
	GetAll(lang string) ([]models.Question, error)
	GetByID(lang, id string) (models.Question, error)

	//LoadQuestions(lang string) ([]models.Question, error)
}

// QuestionServiceStruct is a struct used for passing in the questions file
type ServiceStruct struct {
}

// QuestionService is an instance of QuestionInterface
var Service QuestionInterface = new(ServiceStruct)
