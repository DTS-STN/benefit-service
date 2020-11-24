package renderings

import "github.com/DTS-STN/benefit-service/models"

// BenefitQuestionsResponse is the response containing the information for the requested
// benefits including question dependencies.
type BenefitQuestionsResponse struct {
	// Date period for request in ms since epoch
	RequestDate int64 `json:"request_date"`
	// Array of life journeys, which represent a subset of benefits
	LifeJourneys []string `json:"life_journeys"`
	// Array of specific benefits to get the questions for
	BenefitList []string `json:"benefit_list"`
	// List of answered priority questions
	QuestionList []models.Question `json:"question_list"`
}
