package bindings

import "github.com/DTS-STN/question-priority-service/models"

// BenefitQuestionsRequest is the request sent by the client that contains the benefits
// for which they would like the list of required questions.
type BenefitQuestionsRequest struct {
	// Date period for request in ms since epoch
	RequestDate int64 `json:"request_date"`
	// Array of life journeys, which represent a subset of benefits
	LifeJourneys []string `json:"life_journeys"`
	// Array of specific benefits to get the questions for
	BenefitList []string `json:"benefit_list"`
	// List of answered priority questions
	QuestionList []models.Question `json:"question_list"`
}
