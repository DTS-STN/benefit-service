package handlers

import (
	"strings"
	"github.com/DTS-STN/benefit-service/src/questions"
	"github.com/DTS-STN/benefit-service/src/openfisca"
	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/renderings"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Persons struct {
	Person interface{} `json:"person"`
}

// Method for finding if array contains string
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// GetQuestions ...
// @Summary Returns question dependencies for requested benefits
// @Description Returns question dependencies for requested benefits
// @ID GetBenefitQuestions
// @Accept  json
// @Produce json
// @Success 200 {object} renderings.BenefitQuestionsResponse
// @Failure 400 {object} renderings.BenefitServiceError
// @Failure 404 {object} renderings.BenefitServiceError
// @Failure 500 {object} renderings.BenefitServiceError
// @Router /getquestions [get]
func (h *Handler) GetQuestions(c echo.Context) error {
	// Initialize request and response objects
	getQuestionsRequest := new(bindings.BenefitQuestionsRequest)
	getQuestionsResponse := new(renderings.BenefitQuestionsResponse)

	// Bind incoming data to request object
	if err := c.Bind(getQuestionsRequest); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// Copy request field values to response
	getQuestionsResponse.RequestDate = getQuestionsRequest.RequestDate
	getQuestionsResponse.LifeJourneys = getQuestionsRequest.LifeJourneys
	getQuestionsResponse.BenefitList = getQuestionsRequest.BenefitList

	// Convert UNIX timestamp to required OpenFisca time value format
	t := time.Unix(getQuestionsRequest.RequestDate, 0)
	tstring := t.Format("2006-01-02")

	// Initialize top level "Persons" map
	persons := map[string]interface{}{}

	// Initialize second level "Person" map
	p := map[string]map[string]interface{}{}

	// Get benefits from client request and format for request to OpenFisca
	for _, benefit := range getQuestionsRequest.BenefitList {
		p[benefit] = map[string]interface{}{}
		p[benefit][tstring] = nil
	}
	persons["persons"] = Persons{p}

	// Send request to OpenFisca trace endpoint to find dependencies of requested benefits
	ofResponse, err := openfisca.OFService.SendRequest(openfisca.OFService{}, persons)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// Get trace object out of the OpenFisca response
	trace := ofResponse["trace"].(map[string]interface{})

	// Load questions from local file
	questionList, err := questions.Service.LoadQuestions()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	
	// Loop through OpenFisca trace response
	for benefitKey, benefitObject := range trace {
		// Find benefit keys
		if strings.Contains(benefitKey, "is_eligible") {
			// Remove date from benefit key string
			benefitKey = benefitKey[:strings.IndexByte(benefitKey, '<')]
			// Loop through benefit's child object
			for benefitInfoKey, benefitInfo := range benefitObject.(map[string]interface{}) {
				// Find array of benefit dependencies
				if strings.Contains(benefitInfoKey, "dependencies") {
					// For each dependency in array
					for _, dependency := range benefitInfo.([]interface{}) {
						// Loop through Question List
						for key, question := range questionList {
							// Find question that references the current dependency of the current benefit
							if contains(question.OpenFiscaIds, dependency.(string)[:strings.IndexByte(dependency.(string), '<')]) {
								// Append the current benefit to the DependencyOf array for the current Question
								question.DependencyOf = append(question.DependencyOf, benefitKey)
								questionList[key] = question
							}
						}
					}
				}
			}
		}
	}

	// Add full question list to response to be sent to the client
	getQuestionsResponse.QuestionList = questionList

	return c.JSON(http.StatusOK, getQuestionsResponse)
}
