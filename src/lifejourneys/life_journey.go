package lifejourneys

import (
	"github.com/DTS-STN/benefit-service/models"
)

type LifeJourneyInterface interface {
	LifeJourneys(lang string) []models.LifeJourney
	LifeJourney(lang, id string) (models.LifeJourney, error)
	LoadLifeJourneys(lang string) (lifeJourneys []models.LifeJourney, err error)
}

type LifeJourneyServiceStruct struct {
}

var LifeJourneyService LifeJourneyInterface = new(LifeJourneyServiceStruct)
