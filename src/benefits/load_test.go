package benefits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBenefits(t *testing.T) {
	teardownTests := setupBenefitTests()
	defer teardownTests()

	count := Service.Count()

	assert.Equal(t, 5, count)
}
