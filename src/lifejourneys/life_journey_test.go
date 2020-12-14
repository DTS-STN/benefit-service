package lifejourneys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLifeJourneyBenefitIds(t *testing.T) {
	LifeJourneyService = LifeJourneyServiceStruct{Filename: "../../life_journeys_en.json"}
	lifeJourneyId := "1"
	lifeJourney, err := LifeJourneyService.GetLifeJourneyBenefitById(lifeJourneyId)
	if err != nil {
		assert.Fail(t, "Error occured when getting life journey list")
	}
	assert.Equal(t, lifeJourneyId, lifeJourney.ID)
}

func TestGetAllLifeJourneyBenefits(t *testing.T) {
	LifeJourneyService = LifeJourneyServiceStruct{Filename: "../../life_journeys_en.json"}
	lifeJourney, err := LifeJourneyService.GetAllLifeJourneyBenefits()
	if err != nil {
		assert.Fail(t, "Error occured when getting life journey list")
	}
	assert.Equal(t, 4, len(lifeJourney))
}
