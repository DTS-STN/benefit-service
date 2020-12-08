package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DTS-STN/benefit-service/src/lifejourneys"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLifeJourney(t *testing.T) {
	// Setup Echo service
	e := echo.New()

	lifejourneys.LifeJourneyService = lifejourneys.LifeJourneyServiceStruct{Filename: "../life_journeys_en.json"}
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/lifejourney", nil)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.LifeJourney(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
