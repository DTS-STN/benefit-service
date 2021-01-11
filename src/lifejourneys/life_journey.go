package lifejourneys

import (
	"github.com/DTS-STN/benefit-service/models"
)

type LifeJourneyInterface interface {
	GetAll() ([]models.LifeJourney, error)
	GetByID(id string) (models.LifeJourney, error)
}

type ServiceStruct struct {
	repository LifeJourneyInterface
}

// GetAll returns all Life Journeys
func (q *ServiceStruct) GetAll() ([]models.LifeJourney, error) {
	return q.repository.GetAll()
}

// GetByID returns a Life Journey from an ID
func (q *ServiceStruct) GetByID(id string) (models.LifeJourney, error) {
	return q.repository.GetByID(id)
}

func NewService(repository LifeJourneyRepo) *ServiceStruct {
	return &ServiceStruct{repository: repository}
}

var Service LifeJourneyInterface
