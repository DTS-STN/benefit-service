package lifejourneys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBenefitIds(t *testing.T) {
	Service = ServiceStruct{Filename: "../../life_journeys_en.json"}
	lifeJourneyId := "1"
	lifeJourney, err := Service.GetBenefitById(lifeJourneyId)
	if err != nil {
		assert.Fail(t, "Error occured when getting life journey list")
	}
	assert.Equal(t, lifeJourneyId, lifeJourney.ID)
}

func TestGetAllBenefits(t *testing.T) {
	Service = ServiceStruct{Filename: "../../life_journeys_en.json"}
	lifeJourney, err := Service.GetAllBenefits()
	if err != nil {
		assert.Fail(t, "Error occured when getting life journey list")
	}
	assert.Equal(t, 4, len(lifeJourney))
}
