package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DTS-STN/benefit-service/models"
	"github.com/DTS-STN/benefit-service/src/lifejourneys"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLifeJourney(t *testing.T) {
	// Setup Echo service
	e := echo.New()

	lifejourneys.LifeJourneyService = lifejourneys.LifeJourneyServiceStruct{Filename: "../life_journeys_en.json"}
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/lifejourneys", nil)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)
	q := req.URL.Query()
	q.Add("id", "1")
	// Assertions
	if assert.NoError(t, HandlerService.LifeJourney(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLifeJourneyBenefits(t *testing.T) {
	// Setup Echo service
	e := echo.New()

	lifejourneys.LifeJourneyService = lifejourneys.LifeJourneyServiceStruct{Filename: "../life_journeys_en.json"}
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/lifejourneys/1/benefits", nil)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, HandlerService.LifeJourneyBenefits(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLifeJourneyBenefits_MultipleBenefits(t *testing.T) {
	// Setup Echo service
	e := echo.New()

	lifejourneys.LifeJourneyService = lifejourneys.LifeJourneyServiceStruct{Filename: "../life_journeys_en.json"}
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/lifejourneys/1/benefits", nil)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, HandlerService.LifeJourneyBenefits(c)) {
		benefits := new([]models.Benefits)
		json.NewDecoder(rec.Body).Decode(benefits)

		assert.Equal(t, 3, len(*benefits))
	}
}

func TestGetLifeJourneyBenefitIds(t *testing.T) {
	lifeJourneyId := "1"
	lifeJourney, err := getLifeJourneyBenefitById(lifeJourneyId)
	if err != nil {
		assert.Fail(t, "Error occured when getting life journey list")
	}
	assert.Equal(t, lifeJourneyId, lifeJourney.ID)
}

func TestGetBenefitsByIds(t *testing.T) {
	benefitId := "1"
	benefit, err := getBenefitById(benefitId)
	if err != nil {
		assert.Fail(t, "Error occured when getting benefits by id")
	}
	assert.Equal(t, benefitId, benefit.ID)
}
