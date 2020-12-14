package lifejourneys

import (
	"github.com/DTS-STN/benefit-service/models"
)

type LifeJourneyInterface interface {
	LifeJourneys() []models.LifeJourney
	LifeJourney(id string) (models.LifeJourney, error)
	LoadLifeJourneys() (lifeJourneys []models.LifeJourney, err error)
	ClearLifeJourneys()
	SetFilePath(path string)
}

type LifeJourneyServiceStruct struct {
	Filename string
}

var LifeJourneyService LifeJourneyInterface = new(LifeJourneyServiceStruct)
