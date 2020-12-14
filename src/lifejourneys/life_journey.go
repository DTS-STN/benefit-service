package lifejourneys

import (
	"github.com/DTS-STN/benefit-service/models"
)

type LifeJourneyInterface interface {
	LifeJourneys() []models.LifeJourney
	LoadLifeJourneys() (lifeJourneys []models.LifeJourney, err error)
	GetLifeJourneyBenefitById(id string) (lifeJourney models.LifeJourney, err error)
	GetAllLifeJourneyBenefits() (lifeJourneyList []models.LifeJourney, err error)
}

type LifeJourneyServiceStruct struct {
	Filename string
}

var LifeJourneyService LifeJourneyInterface

func (q LifeJourneyServiceStruct) GetLifeJourneyBenefitById(id string) (lifeJourney models.LifeJourney, err error) {

	ljList, err := LifeJourneyService.LoadLifeJourneys()
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

func (q LifeJourneyServiceStruct) GetAllLifeJourneyBenefits() (lifeJourneyList []models.LifeJourney, err error) {

	ljList, err := LifeJourneyService.LoadLifeJourneys()
	if err != nil {
		return lifeJourneyList, err
	}

	return ljList, nil
}
