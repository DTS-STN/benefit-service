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
	patternMap := make(map[string]interface{})

	//Determine if individual qualifies for Regular EI Benefit
	patternMap["incomeDetails"] = []string{"HFPIR2", "HFPIR1"}
	patternMap["outOfWork"] = []string{"HFPOOW1", "HFPOOW2", "HFPOOW3"}
	patternMap["ableToWork"] = "yes"
	patternMap["reasonForSeparation"] = "HFPRE1"
	patternMap["gender"] = []string{"male", "female"}
	patternMap["province"] = []string{"NL", "PE", "NS", "NB", "QC", "ON", "MB", "SK", "AB", "BC", "YT", "NT", "NU"}
	if benefits.Service.Match(requestMap, patternMap) {
		idArr = append(idArr, 1)
	}
	
	// clear map
	patternMap = make(map[string]interface{})
	//Determine if individual qualifies for Maternity Benefit
	patternMap["incomeDetails"] = []string{"HFPIR2", "HFPIR1"}
	patternMap["outOfWork"] = "HFPOOW1"
	patternMap["ableToWork"] = "no"
	patternMap["reasonForSeparation"] = "HFPRE3"
	patternMap["gender"] = "female"
	patternMap["province"] = []string{"NL", "PE", "NS", "NB", "QC", "ON", "MB", "SK", "AB", "BC", "YT", "NT", "NU"}
	if benefits.Service.Match(requestMap, patternMap) {
		idArr = append(idArr, 2)
	}

	// clear map
	patternMap = make(map[string]interface{})
	//Determine if individual qualifies for Sickness Benefit
	patternMap["incomeDetails"] = []string{"HFPIR2", "HFPIR1"}
	patternMap["outOfWork"] = []string{"HFPOOW1", "HFPOOW2"}
	patternMap["ableToWork"] = "no"
	patternMap["reasonForSeparation"] = "HFPRE2"
	patternMap["gender"] = []string{"male", "female"}
	patternMap["province"] = []string{"NL", "PE", "NS", "NB", "QC", "ON", "MB", "SK", "AB", "BC", "YT", "NT", "NU"}
	if benefits.Service.Match(requestMap, patternMap) {
		idArr = append(idArr, 3)
	}

	// clear map
	patternMap = make(map[string]interface{})
	// check Ontario child care subsidy
	patternMap["reasonForSeparation"] = "HFPRE3"
	patternMap["province"] = "ON"
	if benefits.Service.Match(requestMap, patternMap) {
		idArr = append(idArr, 4)
	}

	// clear map
	patternMap = make(map[string]interface{})
	// check Ontario low income credit
	patternMap["incomeDetails"] = "HFPIR3"
	patternMap["province"] = "ON"
	if benefits.Service.Match(requestMap, patternMap) {
		idArr = append(idArr, 5)
	}

	return c.JSON(http.StatusOK, idArr)
}
