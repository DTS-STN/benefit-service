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

func setupLifeJourneyTests() func() {
	lifejourneys.Files = map[string]string{
		"en": "../life_journeys_en.json",
		"fr": "../life_journeys_fr.json",
	}
	return func() {
	}
}

type LifeJourneyServiceMock struct {
	mock.Mock
}

func (q *LifeJourneyServiceMock) LifeJourneys(lang string) []models.LifeJourney {
	return []models.LifeJourney{
		{
			ID: "1",
		},
	}
}

func (q *LifeJourneyServiceMock) LoadLifeJourneys(lang string) ([]models.LifeJourney, error) {
	return []models.LifeJourney{}, nil
}

func (q *LifeJourneyServiceMock) GetById(lang, id string) (models.LifeJourney, error) {
	return models.LifeJourney{ID: "2", RelatedBenefits: []string{"1", "2", "3"}}, nil
}

func TestLifeJourney(t *testing.T) {
	teardownTests := setupLifeJourneyTests()
	defer teardownTests()

	// Setup Echo service
	e := echo.New()

	lifejourneys.Service = lifejourneys.LifeJourneyInterface(new(LifeJourneyServiceMock))
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/lifejourneys", nil)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)
	q := req.URL.Query()
	q.Add("id", "1")
	q.Add("lang", "en")
	// Assertions
	if assert.NoError(t, HandlerService.LifeJourneys(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLifeJourneyBenefits(t *testing.T) {
	teardownTests := setupLifeJourneyTests()
	defer teardownTests()

	// Setup Echo service
	e := echo.New()

	lifejourneys.Service = lifejourneys.LifeJourneyInterface(new(LifeJourneyServiceMock))
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
	teardownTests := setupLifeJourneyTests()
	defer teardownTests()

	// Setup Echo service
	e := echo.New()

	lifejourneys.Service = lifejourneys.LifeJourneyInterface(new(LifeJourneyServiceMock))
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
