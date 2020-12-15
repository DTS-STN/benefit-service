package benefits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBenefitsByIds(t *testing.T) {
	Service = ServiceStruct{Filename: "../../benefit_info_en.json"}
	benefitId := "1"
	benefit, err := Service.GetBenefitById(benefitId)
	if err != nil {
		assert.Fail(t, "Error occured when getting benefits by id")
	}
	assert.Equal(t, benefitId, benefit.ID)
}
