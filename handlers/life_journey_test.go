package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DTS-STN/benefit-service/renderings"
	"github.com/DTS-STN/benefit-service/src/lifejourneys"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLifeJourney(t *testing.T) {
	// Setup Echo service
	e := echo.New()

	lifejourneys.LifeJourneyService = lifejourneys.LifeJourneyServiceStruct{Filename: "../life_journeys.json"}
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

func TestLifeJourney_AllLifeJourneys(t *testing.T) {
	// Setup Echo service
	e := echo.New()

	lifejourneys.LifeJourneyService = lifejourneys.LifeJourneyServiceStruct{Filename: "../life_journeys.json"}
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/lifejourney", nil)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.LifeJourney(c)) {
		lifeJourneyResponse := new(renderings.LifeJourneyResponse)
		json.NewDecoder(rec.Body).Decode(lifeJourneyResponse)

		assert.Equal(t, 4, len(lifeJourneyResponse.LifeJourneyList))
	}
}

func TestLifeJourney_SingleLifeJourney(t *testing.T) {
	// Setup Echo service
	e := echo.New()

	lifejourneys.LifeJourneyService = lifejourneys.LifeJourneyServiceStruct{Filename: "../life_journeys.json"}
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/lifejourney", nil)
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.LifeJourney(c)) {
		lifeJourneyResponse := new(renderings.LifeJourneyResponse)
		json.NewDecoder(rec.Body).Decode(lifeJourneyResponse)

		assert.Equal(t, 1, len(lifeJourneyResponse.LifeJourneyList))
	}
}
