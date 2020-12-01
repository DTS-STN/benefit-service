package renderings

import "github.com/DTS-STN/benefit-service/models"

// LifeJourneyResponse is the response returned to the client that contains
// information on Life journeys and a list of related lifejourneys
type LifeJourneyResponse struct {
	// Life Journey ID
	LifeJourneyList []models.LifeJourney `json:"lifejourneys"`
}
