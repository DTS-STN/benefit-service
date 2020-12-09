package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DTS-STN/benefit-service/renderings"
	"github.com/DTS-STN/benefit-service/src/benefits"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBenefits(t *testing.T) {
	// Setup Echo service
	e := echo.New()
	//Load Benefits Data File
	benefits.BenefitsService = &benefits.BenefitsServiceStruct{Filename: "../benefit_info_en.json"}
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/benefits", nil)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.Benefits(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestBenefits_AllBenefits(t *testing.T) {
	// Setup Echo service
	e := echo.New()
	//Load Benefits Data File
	benefits.BenefitsService = &benefits.BenefitsServiceStruct{Filename: "../benefit_info_en.json"}
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/benefits", nil)

	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.Benefits(c)) {
		benefitResponse := new(renderings.BenefitsResponse)
		json.NewDecoder(rec.Body).Decode(&benefitResponse.BenefitsList)

		assert.Equal(t, 5, len(benefitResponse.BenefitsList))
	}
}

func TestBenefits_SingleBenefit(t *testing.T) {
	// Setup Echo service
	e := echo.New()
	//Load Benefits Data File
	benefits.BenefitsService = &benefits.BenefitsServiceStruct{Filename: "../benefit_info_en.json"}
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/benefits", nil)
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()

	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.Benefits(c)) {
		benefitResponse := new(renderings.BenefitsResponse)
		json.NewDecoder(rec.Body).Decode(&benefitResponse.BenefitsList)

		assert.Equal(t, 1, len(benefitResponse.BenefitsList))
	}
}

func TestBenefits_MultipleBenefits(t *testing.T) {
	// Setup Echo service
	e := echo.New()
	//Load Benefits Data File
	benefits.BenefitsService = &benefits.BenefitsServiceStruct{Filename: "../benefit_info_en.json"}
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/benefits", nil)
	q := req.URL.Query()
	q.Add("idList", "1,2,3")
	req.URL.RawQuery = q.Encode()

	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.Benefits(c)) {
		benefitResponse := new(renderings.BenefitsResponse)
		json.NewDecoder(rec.Body).Decode(&benefitResponse.BenefitsList)

		assert.Equal(t, 3, len(benefitResponse.BenefitsList))
	}
}
