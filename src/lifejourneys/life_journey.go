package lifejourneys

import (
	"github.com/DTS-STN/benefit-service/models"
)

type LifeJourneyInterface interface {
	GetAll(lang string) []models.LifeJourney
	GetByID(lang, id string) (models.LifeJourney, error)
	LoadLifeJourneys(lang string) ([]models.LifeJourney, error)
}

type ServiceStruct struct {
}

var Service LifeJourneyInterface = new(ServiceStruct)
