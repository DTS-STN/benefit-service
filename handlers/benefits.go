package handlers

import (
	"net/http"
	"strings"

	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/renderings"
	"github.com/DTS-STN/benefit-service/src/benefits"
	"github.com/labstack/echo/v4"
)

// Benefits
// @Summary Request Info on Benefits
// @ID benefits
// @Accept  json
// @Produce json
// @Success 200 {object} renderings.BenefitsResponse
// @Failure 400 {object} renderings.BenefitServiceError
// @Failure 404 {object} renderings.BenefitServiceError
// @Failure 500 {object} renderings.BenefitServiceError
// @Param lang query string false "The language the response should be in. Defaults to English. English and French supported."
// @Router /benefits [get]
func (h *Handler) Benefits(c echo.Context) error {
	var benefitsResponse = new(renderings.BenefitsResponse)
	benefitsRequest := new(bindings.BenefitsRequest)

	// Bind the request into our request struct
	if err := c.Bind(benefitsRequest); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, benefitsResponse)
	}

	setLanguage(benefitsRequest.Lang)
	benList, err := benefits.BenefitsService.LoadBenefits()

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, benefitsResponse)
	}

	if benefitsRequest.Id != "" {
		for _, ben := range benList {
			if ben.ID == benefitsRequest.Id {
				benefitsResponse.BenefitsList = append(benefitsResponse.BenefitsList, ben)
				break
			}

		}
	} else if benefitsRequest.IdList != "" {
		Ids := strings.Split(benefitsRequest.IdList, ",")
		for _, benId := range Ids {
			for _, ben := range benList {
				if ben.ID == benId {
					benefitsResponse.BenefitsList = append(benefitsResponse.BenefitsList, ben)
					break
				}
			}
		}
	} else {
		benefitsResponse.BenefitsList = benList
	}
	return c.JSON(http.StatusOK, benefitsResponse.BenefitsList)
}
