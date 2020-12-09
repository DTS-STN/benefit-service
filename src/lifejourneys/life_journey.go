package lifejourneys

import (
	"github.com/DTS-STN/benefit-service/models"
)

type LifeJourneyInterface interface {
	LifeJourneys() []models.LifeJourney
	LoadLifeJourneys() (lifeJourneys []models.LifeJourney, err error)
	SetFilePath(path string)
}

type LifeJourneyServiceStruct struct {
	Filename string
}

var LifeJourneyService LifeJourneyInterface = new(LifeJourneyServiceStruct)
