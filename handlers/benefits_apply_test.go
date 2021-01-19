package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DTS-STN/benefit-service/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBenefitsApply(t *testing.T) {
	// Setup Echo service
	e := echo.New()

	// Setup http request using httptest
	benefit := new(models.Benefits)
	benefit.BenefitType = "something"
	benefit_json, _ := json.Marshal(benefit)
	benefit_json_string_reader := strings.NewReader(string(benefit_json))
	req := httptest.NewRequest(http.MethodPost, "/benefits/apply", benefit_json_string_reader)
	req.Header.Set("Content-Type", "application/json")

	// Setup response
	rec := httptest.NewRecorder()

	// Create a new Echo Context
	c := e.NewContext(req, rec)

	if assert.NoError(t, HandlerService.BenefitsApply(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

}
