package lifejourneys

import (
	"github.com/DTS-STN/benefit-service/models"
)

type LifeJourneyInterface interface {
	LifeJourneys(lang string) []models.LifeJourney
	GetById(lang, id string) (models.LifeJourney, error)
	LoadLifeJourneys(lang string) ([]models.LifeJourney, error)
}

type ServiceStruct struct {
}

var Service LifeJourneyInterface = new(ServiceStruct)
