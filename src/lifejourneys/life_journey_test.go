package lifejourneys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupBenefitTests() func() {
	Files = map[string]string{
		"en": "../../life_journeys_en.json",
		"fr": "../../life_journeys_fr.json",
	}
	return func() {

	}
}

func TestGetLifeJourenyIds(t *testing.T) {
	teardownTests := setupBenefitTests()
	defer teardownTests()

	lifeJourneyId := "1"
	lang := "en"
	lifeJourney, err := Service.GetByID(lang, lifeJourneyId)
	if err != nil {
		assert.Fail(t, "Error occured when getting life journey list")
	}
	assert.Equal(t, lifeJourneyId, lifeJourney.ID)
}

func TestGetAllLifeJourneys(t *testing.T) {
	teardownTests := setupBenefitTests()
	defer teardownTests()

	lang := "en"
	lifeJourney := Service.GetAll(lang)

	assert.Equal(t, 4, len(lifeJourney))
}
