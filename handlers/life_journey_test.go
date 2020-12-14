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
	"github.com/stretchr/testify/mock"
)

type LifeJourneyServiceMock struct {
	mock.Mock
}

func (q *LifeJourneyServiceMock) LifeJourneys() (lifeJourneys []models.LifeJourney) {
	return lifeJourneys
}

func (q *LifeJourneyServiceMock) LoadLifeJourneys() (lifeJourneys []models.LifeJourney, err error) {
	return lifeJourneys, nil
}

func (q *LifeJourneyServiceMock) GetLifeJourneyBenefitById(id string) (lifeJourney models.LifeJourney, err error) {
	lifeJourney.ID = "2"
	lifeJourney.RelatedBenefits = []string{"1", "2", "3"}
	return lifeJourney, nil
}

func (q *LifeJourneyServiceMock) GetAllLifeJourneyBenefits() (lifeJourneyList []models.LifeJourney, err error) {
	lifeJourney := new(models.LifeJourney)
	lifeJourney.ID = "1"
	lifeJourneyList = append(lifeJourneyList, *lifeJourney)
	return lifeJourneyList, nil
}

func TestLifeJourney(t *testing.T) {
	// Setup Echo service
	e := echo.New()

	lifejourneys.LifeJourneyService = lifejourneys.LifeJourneyInterface(new(LifeJourneyServiceMock))
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

	lifejourneys.LifeJourneyService = lifejourneys.LifeJourneyInterface(new(LifeJourneyServiceMock))
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

	lifejourneys.LifeJourneyService = lifejourneys.LifeJourneyInterface(new(LifeJourneyServiceMock))
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
