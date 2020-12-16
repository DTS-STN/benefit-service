package benefits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupBenefitTests() func() {
	Files = map[string]string{
		"en": "../../benefit_info_en.json",
		"fr": "../../benefit_info_fr.json",
	}
	return func() {

	}
}

func TestGetBenefitsByIds(t *testing.T) {
	teardownTests := setupBenefitTests()
	defer teardownTests()

	benefitId := "1"
	lang := "en"
	benefit, err := Service.GetByID(lang, benefitId)
	if err != nil {
		assert.Fail(t, "Error occured when getting benefits by id")
	}
	assert.Equal(t, benefitId, benefit.ID)
}
