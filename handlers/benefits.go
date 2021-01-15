package handlers

import (
	"net/http"
	"strings"

	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/models"
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
	var err error

	// Bind the request into our request struct
	if err = c.Bind(benefitsRequest); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, benefitsResponse)
	}

	// if an ID is passed in, get the benefit based on the ID
	if benefitsRequest.Id != "" {
		if benefitsResponse.Benefit, err = benefits.Service.GetByID(benefitsRequest.Lang, benefitsRequest.Id); err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, benefitsResponse)
		}

		return c.JSON(http.StatusOK, benefitsResponse.Benefit)
	}

	// if an ID list is passed in, get all the benefits for the IDs
	if benefitsRequest.IdList != "" {
		Ids := strings.Split(benefitsRequest.IdList, ",")
		var benefit models.Benefits
		for _, benId := range Ids {
			if benefit, err = benefits.Service.GetByID(benefitsRequest.Lang, benId); err != nil {
				c.Logger().Error(err)
				continue
			}
			benefitsResponse.BenefitsList = append(benefitsResponse.BenefitsList, benefit)
		}
		return c.JSON(http.StatusOK, benefitsResponse.BenefitsList)
	}

	// otherwise return all benefits
	benefitsResponse.BenefitsList = benefits.Service.GetAll(benefitsRequest.Lang)
	return c.JSON(http.StatusOK, benefitsResponse.BenefitsList)
}

func (h *Handler) BenefitsCount(c echo.Context) error {
	count := benefits.Service.Count()
	return c.JSON(http.StatusOK, count)
}
