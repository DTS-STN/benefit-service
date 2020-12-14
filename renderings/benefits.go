package renderings

import "github.com/DTS-STN/benefit-service/models"

// BenefitsResponse is the response returned to the client that contains
// information on Benefits
type BenefitsResponse struct {
	// Life Journey ID
	BenefitsList []models.Benefits `json:"benefits"`
	Benefit      models.Benefits
}
