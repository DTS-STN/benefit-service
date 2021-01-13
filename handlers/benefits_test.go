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

		assert.Equal(t, 3, len(benefitResponse.BenefitsList))
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
			Title:           "Employment Insurance",
			Description:     "Employment Insurance Lorem Ipsum Dolor Sit Amet.",
			LongDescription: "# Employment Insurance Details \n* Lorem ipsum dolor sit amet \n## Curabitur feugiat, turpis a dignissim dictum \n* Praesent fermentum lectus ac vulputate suscipit  \n## Aliquam vehicula consectetur felis ac luctus \n* Praesent et sollicitudin felis, vitae lobortis sapien \n## Pellentesque consequat \n* Suspendisse ac posuere tortor, consequat imperdiet augue \n* Must have a permanant address \n* Must be a resident of Canada \n* Must be a Canadian Citizen \n### Vestibulum mollis in dolor in pretium \n* Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos \n## Phasellus varius \n* Pellentesque consequat diam in rhoncus dapibus \n## Quisque tempus \n* Morbi sit amet varius mi, ut viverra lorem. \n## Duis vestibulum \n* Cras fringilla euismod ante sit amet consequat. \n## Donec rutrum \n* Donec ultrices ultricies ipsum, ut iaculis sapien euismod eget.",
			RelatedBenefits: []string{},
			BenefitDetails: []models.FieldDetails{
				{
					FieldName:             "Overview",
					FieldShortDescription: "Short Overview of Employment Insurance Process",
					FieldLongDescription:  "Long Overview of Employment Insurance Process",
				},
				{
					FieldName:             "Eligibility criteria",
					FieldShortDescription: "Short description of Employment Insurance Eligibility Criteria",
					FieldLongDescription:  "Long description of Employment Insurance Eligibility Criteria",
				},
				{
					FieldName:             "Eligibility period",
					FieldShortDescription: "Short description of Eligibility periods for Employment Insurance",
					FieldLongDescription:  "Long description of Eligibility periods for Employment Insurance",
				},
				{
					FieldName:             "Important Information",
					FieldShortDescription: "Short description of Important Information for Employment Insurance",
					FieldLongDescription:  "Long description of Important Information for Employment Insurance",
				},
				{
					FieldName:             "How to apply",
					FieldShortDescription: "Short description of How to apply for Employment Insurance",
					FieldLongDescription:  "Long description of How to apply for Employment Insurance",
				},
				{
					FieldName:             "Contact Information",
					FieldShortDescription: "Short description of Contact Information for Employment Insurance",
					FieldLongDescription:  "Long description of Contact Information for Employment Insurance",
				},
				{
					FieldName:             "Examples",
					FieldShortDescription: "Short description of examples for Employment Insurance",
					FieldLongDescription:  "Long description of examples for Employment Insurance",
				},
				{
					FieldName:             "Payment Information",
					FieldShortDescription: "Short description of Payment Information for Employment Insurance",
					FieldLongDescription:  "Long description of Payment Information for Employment Insurance",
				},
				{
					FieldName:             "Repayment Information",
					FieldShortDescription: "Short description of Repayment Information for Employment Insurance",
					FieldLongDescription:  "Long description of Repayment Information for Employment Insurance",
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
