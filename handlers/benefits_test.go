package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DTS-STN/benefit-service/models"
	"github.com/DTS-STN/benefit-service/renderings"
	"github.com/DTS-STN/benefit-service/src/benefits"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupBenefitTests() func() {
	benefits.Files = map[string]string{
		"en": "../benefit_info_en.json",
		"fr": "../benefit_info_fr.json",
	}
	return func() {

	}
}

func TestBenefits(t *testing.T) {
	teardownTests := setupBenefitTests()
	defer teardownTests()

	// Setup Echo service
	e := echo.New()
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
	teardownTests := setupBenefitTests()
	defer teardownTests()

	// Setup Echo service
	e := echo.New()
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
	teardownTests := setupBenefitTests()
	defer teardownTests()

	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/benefits", nil)
	q := req.URL.Query()
	q.Add("id", "1")
	q.Add("lang", "en")
	req.URL.RawQuery = q.Encode()

	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.Benefits(c)) {
		benefitResponse := new(renderings.BenefitsResponse)
		json.NewDecoder(rec.Body).Decode(&benefitResponse.Benefit)

		expectedResult := models.Benefits{
			ID:              "1",
			Title:           "Driver's License",
			Description:     "License for Driver's",
			LongDescription: "# Driver's License Benefit Details \nBrief description of the Driver's License Benefit \n## Overview \nShort Overview of Driver's License Process \n## Important Information \nImportant Information cli8ents need to know for the Driver's License benefit. \n## Eligibility criteria \nDescription of Driver's License Eligibility Criteria and how to qualify for the benefit. \n- Must be 16 years of age or older \n- Must be a resident of Canada \n- Must have completed a Driver Training Program \n- For client's under the age of 18, you must have parental consent \n### Examples \nDescription of different scenario's to provide examples to clients \n## Eligibility period \nDescription of Eligibility periods for a Driver's License. \n## How to apply \nDescription of How to apply for a Driver's License and what information is required. \n## Contact Information \nFor further Information on Driver's License and related Benefits contact 1-800-Drivers. \n## Payment Information \nDescription of Payment Information for a Driver's License.",
			RelatedBenefits: []string{"5"},
			BenefitDetails: []models.FieldDetails{
				{
					FieldName:             "Overview",
					FieldShortDescription: "Short Overview of Driver's License Process",
					FieldLongDescription:  "Long Overview of Driver's License Process",
				},
				{
					FieldName:             "Eligibility criteria",
					FieldShortDescription: "Short description of Driver's License Eligibility Criteria",
					FieldLongDescription:  "Long description of Driver's License Eligibility Criteria",
				},
				{
					FieldName:             "Eligibility period",
					FieldShortDescription: "Short description of Eligibility periods for a Driver's License",
					FieldLongDescription:  "Long description of Eligibility periods for a Driver's License",
				},
				{
					FieldName:             "Important Information",
					FieldShortDescription: "Short description of Important Information for a Driver's License",
					FieldLongDescription:  "Long description of Important Information for a Driver's License",
				},
				{
					FieldName:             "How to apply",
					FieldShortDescription: "Short description of How to apply for a Driver's License",
					FieldLongDescription:  "Long description of How to apply for a Driver's License",
				},
				{
					FieldName:             "Contact Information",
					FieldShortDescription: "Short description of Contact Information for a Driver's License",
					FieldLongDescription:  "Long description of Contact Information for a Driver's License",
				},
				{
					FieldName:             "Examples",
					FieldShortDescription: "Short description of examples for a Driver's License",
					FieldLongDescription:  "Long description of examples for a Driver's License",
				},
				{
					FieldName:             "Payment Information",
					FieldShortDescription: "Short description of Payment Information for a Driver's License",
					FieldLongDescription:  "Long description of Payment Information for a Driver's License",
				},
				{
					FieldName:             "Repayment Information",
					FieldShortDescription: "Short description of Repayment Information for a Driver's License",
					FieldLongDescription:  "Long description of Repayment Information for a Driver's License",
				},
			},
		}
		assert.Equal(t, expectedResult, benefitResponse.Benefit)
	}
}

func TestBenefits_MultipleBenefits(t *testing.T) {
	teardownTests := setupBenefitTests()
	defer teardownTests()

	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/benefits", nil)
	q := req.URL.Query()
	q.Add("idList", "1,2,3")
	q.Add("lang", "en")
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
