package benefits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBenefits(t *testing.T) {
	teardownTests := setupBenefitTests()
	defer teardownTests()

	lang := "en"
	count := Service.Count(lang)

	assert.Equal(t, 3, count)
}
