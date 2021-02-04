package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/src/benefits"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupEligibleTests() func() {
	benefits.Files = map[string]string{
		"en": "../benefit_info_en.json",
		"fr": "../benefit_info_fr.json",
	}
	return func() {

	}
}

func TestEligible(t *testing.T) {
	teardownTests := setupEligibleTests()
	defer teardownTests()

	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	req := httptest.NewRequest("POST", "/benefits/eligible", nil)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.BenefitsEligibility(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestBenefits_EligibleOnlyRegular(t *testing.T) {
	teardownTests := setupBenefitTests()
	defer teardownTests()

	benefit := bindings.BenefitEligibilityRequest{
		IncomeDetails:      "gt-60k",
		TimeOutOfWork:      "lt-2weeks",
		AbleToWork:         "yes",
		ReasonForOutOfWork: "lost-job",
		Gender:             "male",
	}

	benefit_json, _ := json.Marshal(benefit)
	benefitJSON := string(benefit_json)

	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	req := httptest.NewRequest("POST", "/benefits/eligible", strings.NewReader(benefitJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.BenefitsEligibility(c)) {
		benefitResponse := []int{}
		json.NewDecoder(rec.Body).Decode(&benefitResponse)

		expectedResult := []int{1}

		assert.Equal(t, expectedResult, benefitResponse)
	}
}

func TestBenefits_EligibleOnlyMaternity(t *testing.T) {
	teardownTests := setupBenefitTests()
	defer teardownTests()

	benefit := bindings.BenefitEligibilityRequest{
		IncomeDetails:      "gt-60k",
		TimeOutOfWork:      "lt-2weeks",
		AbleToWork:         "no",
		ReasonForOutOfWork: "baby",
		Gender:             "female",
	}

	benefit_json, _ := json.Marshal(benefit)
	benefitJSON := string(benefit_json)

	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	req := httptest.NewRequest("POST", "/benefits/eligible", strings.NewReader(benefitJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.BenefitsEligibility(c)) {
		benefitResponse := []int{}
		json.NewDecoder(rec.Body).Decode(&benefitResponse)

		expectedResult := []int{2}

		assert.Equal(t, expectedResult, benefitResponse)
	}
}

func TestBenefits_EligibleOnlySickness(t *testing.T) {
	teardownTests := setupBenefitTests()
	defer teardownTests()

	benefit := bindings.BenefitEligibilityRequest{
		IncomeDetails:      "gt-60k",
		TimeOutOfWork:      "lt-2weeks",
		AbleToWork:         "no",
		ReasonForOutOfWork: "sick",
		Gender:             "male",
	}

	benefit_json, _ := json.Marshal(benefit)
	benefitJSON := string(benefit_json)

	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	req := httptest.NewRequest("POST", "/benefits/eligible", strings.NewReader(benefitJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.BenefitsEligibility(c)) {
		benefitResponse := []int{}
		json.NewDecoder(rec.Body).Decode(&benefitResponse)

		expectedResult := []int{3}

		assert.Equal(t, expectedResult, benefitResponse)
	}
}

func TestBenefits_NonEligible(t *testing.T) {
	teardownTests := setupBenefitTests()
	defer teardownTests()

	benefit := bindings.BenefitEligibilityRequest{
		IncomeDetails:      "lt-30k",
		TimeOutOfWork:      "lt-2weeks",
		AbleToWork:         "yes",
		ReasonForOutOfWork: "lost-job",
		Gender:             "male",
	}

	benefit_json, _ := json.Marshal(benefit)
	benefitJSON := string(benefit_json)

	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	req := httptest.NewRequest("POST", "/benefits/eligible", strings.NewReader(benefitJSON))
	req.Header.Set("Content-Type", "application/json")

	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.BenefitsEligibility(c)) {
		benefitResponse := []int{}
		json.NewDecoder(rec.Body).Decode(&benefitResponse)

		expectedResult := []int(nil)

		assert.Equal(t, expectedResult, benefitResponse)
	}
}