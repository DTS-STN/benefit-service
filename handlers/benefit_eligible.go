package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/models"
	"github.com/DTS-STN/benefit-service/src/benefits"
	"github.com/labstack/echo/v4"
)

// Determine benefit eligibility
// @Summary Request ids of eligible benefits
// @ID benefits-eligible
// @Accept  json
// @Produce json
// @Success 200 {array} int
// @Failure 400 {object} models.Error
// @Param requestBody body bindings.BenefitEligibilityRequest true "The answers to the questions that determine benefit eligibility"
// @Router /benefits/eligible [post]
func (h *Handler) BenefitsEligibility(c echo.Context) error {
	var err error
	var requestMap map[string]interface{}
	var idArr []int

	//Bind the request into our request struct
	eligible := new(bindings.BenefitEligibilityRequest)
	if err = c.Bind(eligible); err != nil {
		errObj := new(models.Error)
		errObj.Status = http.StatusBadRequest
		errObj.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, errObj)
	}

	//Decode request into a map
	data, _ := json.Marshal(eligible)
	json.Unmarshal(data, &requestMap)

	//Create a map to compare against
	patternMap := map[string]interface{}{}

	//Determine if individual qualifies for Regular EI Benefit
	patternMap["incomeDetails"] = []string{"30k-to-60k", "gt-60k"}
	patternMap["timeOutOfWork"] = []string{"lt-2weeks", "2weeks-3months", "gt-3months"}
	patternMap["ableToWork"] = "yes"
	patternMap["reasonForOutOfWork"] = "lost-job"
	patternMap["gender"] = []string{"male", "female"}
	if benefits.Service.Match(requestMap, patternMap) {
		idArr = append(idArr, 1)
	}

	//Determine if individual qualifies for Maternity Benefit
	patternMap["incomeDetails"] = []string{"30k-to-60k", "gt-60k"}
	patternMap["timeOutOfWork"] = "lt-2weeks"
	patternMap["ableToWork"] = "no"
	patternMap["reasonForOutOfWork"] = "baby"
	patternMap["gender"] = "female"
	if benefits.Service.Match(requestMap, patternMap) {
		idArr = append(idArr, 2)
	}

	//Determine if individual qualifies for Sickness Benefit
	patternMap["incomeDetails"] = []string{"30k-to-60k", "gt-60k"}
	patternMap["timeOutOfWork"] = []string{"lt-2weeks", "2weeks-3months"}
	patternMap["ableToWork"] = "no"
	patternMap["reasonForOutOfWork"] = "sick"
	patternMap["gender"] = []string{"male", "female"}
	if benefits.Service.Match(requestMap, patternMap) {
		idArr = append(idArr, 3)
	}
	return c.JSON(http.StatusOK, idArr)
}
