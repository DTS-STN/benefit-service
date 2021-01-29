package handlers

import (
	"net/http"

	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/models"
	"github.com/DTS-STN/benefit-service/renderings"
	"github.com/DTS-STN/benefit-service/src/benefits"
	"github.com/labstack/echo/v4"
)

func (h *Handler) BenefitsEligibility(c echo.Context) error {
	var benefitsResponse = new(renderings.BenefitsResponse)
	var benefit models.Benefits
	var err error

	//Bind the request into our request struct
	eligible := new(bindings.BenefitEligibilityRequest)
	if err := c.Bind(eligible); err != nil {
		errObj := new(models.Error)
		errObj.Status = http.StatusBadRequest
		errObj.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, errObj)
	}

	//Check to see if any of the fields are empty
	//TODO: Clean this up/reduce redundant code
	if eligible.IncomeDetails == "" {
		errObj := new(models.Error)
		errObj.Status = http.StatusBadRequest
		errObj.ErrorMessage = "incomeDetails were not supplied"
		return c.JSON(http.StatusBadRequest, errObj)
	}

	if eligible.TimeOutOfWork == "" {
		errObj := new(models.Error)
		errObj.Status = http.StatusBadRequest
		errObj.ErrorMessage = "timeOutOfWork was not supplied"
		return c.JSON(http.StatusBadRequest, errObj)
	}

	if eligible.AbleToWork == "" {
		errObj := new(models.Error)
		errObj.Status = http.StatusBadRequest
		errObj.ErrorMessage = "ableToWork was not supplied"
		return c.JSON(http.StatusBadRequest, errObj)
	}

	if eligible.ReasonForOutOfWork == "" {
		errObj := new(models.Error)
		errObj.Status = http.StatusBadRequest
		errObj.ErrorMessage = "reasonForOutOfWork was not supplied"
		return c.JSON(http.StatusBadRequest, errObj)
	}

	if eligible.Gender == "" {
		errObj := new(models.Error)
		errObj.Status = http.StatusBadRequest
		errObj.ErrorMessage = "gender was not supplied"
		return c.JSON(http.StatusBadRequest, errObj)
	}

	//Determine if individual qualifies for Regular EI Benefit
	if (eligible.IncomeDetails != "lt-30k") &&
		(eligible.AbleToWork == "yes") &&
		(eligible.ReasonForOutOfWork == "lost-job") {
		
		//Get Regular EI Benefit and append it to the BenefitsList
		if benefit, err = benefits.Service.GetByID("en", "1"); err != nil {
			c.Logger().Error(err)
		}
		benefitsResponse.BenefitsList = append(benefitsResponse.BenefitsList, benefit)
	}

	//Determine if individual qualifies for Maternity Benefit
	if (eligible.IncomeDetails != "lt-30k") &&
		(eligible.TimeOutOfWork == "lt-2weeks") &&
		(eligible.AbleToWork == "no") &&
		(eligible.ReasonForOutOfWork == "baby") &&
		(eligible.Gender == "female") {

		//Get Maternity benefit	and append it to the BenefitsList
		if benefit, err = benefits.Service.GetByID("en", "2"); err != nil {
			c.Logger().Error(err)
		}
		benefitsResponse.BenefitsList = append(benefitsResponse.BenefitsList, benefit)
	}

	//Determine if individual qualifies for Sickness Benefit
	if (eligible.IncomeDetails != "lt-30k") &&
		(eligible.TimeOutOfWork != "gt-3months") &&
		(eligible.AbleToWork == "no") &&
		(eligible.ReasonForOutOfWork == "sick") {

		//Get sickness benefit and append it to the BenefitsList	
		if benefit, err = benefits.Service.GetByID("en", "3"); err != nil {
			c.Logger().Error(err)
		}
		benefitsResponse.BenefitsList = append(benefitsResponse.BenefitsList, benefit)
	}

	return c.JSON(http.StatusOK, benefitsResponse.BenefitsList)
}
