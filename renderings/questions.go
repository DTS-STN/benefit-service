package renderings

import "github.com/DTS-STN/benefit-service/models"

// QuestionResponse is the response returned to the client that contains
// information on questinos
type QuestionResponse struct {
	QuestionList []models.Question `json:"questions"`
	Question     models.Question
}
