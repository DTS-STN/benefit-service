package lifejourneys

import (
	"github.com/DTS-STN/benefit-service/models"
)

type LifeJourneyInterface interface {
	LifeJourneys() []models.LifeJourney
	LoadLifeJourneys() (lifeJourneys []models.LifeJourney, err error)
	GetById(id string) (lifeJourney models.LifeJourney, err error)
	GetAllBenefits() (lifeJourneyList []models.LifeJourney, err error)
}

type ServiceStruct struct {
	Filename string
}

var Service LifeJourneyInterface

func (q ServiceStruct) GetById(id string) (lifeJourney models.LifeJourney, err error) {

	ljList, err := Service.LoadLifeJourneys()
	if err != nil {
		return lifeJourney, err
	}

	for _, lj := range ljList {
		if lj.ID == id {
			lifeJourney = lj
			break
		}
	}

	return lifeJourney, nil
}

func (q ServiceStruct) GetAllBenefits() (lifeJourneyList []models.LifeJourney, err error) {

	ljList, err := Service.LoadLifeJourneys()
	if err != nil {
		return lifeJourneyList, err
	}

	return ljList, nil
}
