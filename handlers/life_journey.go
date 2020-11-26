package handlers

import (
	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/renderings"
	"github.com/DTS-STN/benefit-service/src/lifejourneys"
	"github.com/labstack/echo/v4"
	"net/http"
)

// LifeJourney
// @Summary Request Info on Life Journey
// @ID Life Journey
// @Accept  json
// @Produce json
// @Success 200 {object} renderings.LifeJourneyResponse
// @Failure 400 {object} renderings.BenefitServiceError
// @Failure 404 {object} renderings.BenefitServiceError
// @Failure 500 {object} renderings.BenefitServiceError
// @Router /lifejourney [get]

func (h *Handler) LifeJourney(c echo.Context) error {
	var lifeJourneyResponse = new(renderings.LifeJourneyResponse)
	lifeJourneyRequest := new(bindings.LifeJourneyRequest)

	// Bind the request into our request struct
	if err := c.Bind(lifeJourneyRequest); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, lifeJourneyResponse)
	}
	ljList, err := lifejourneys.LifeJourneyService.LoadLifeJourneys()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, lifeJourneyResponse)
	}

	if lifeJourneyRequest.Id != 0 {
		for _, lj := range ljList {
			if lj.ID == lifeJourneyRequest.Id {
				lifeJourneyResponse.LifeJourneyList = append(lifeJourneyResponse.LifeJourneyList, lj)
				break
			}

		}
	} else {
		lifeJourneyResponse.LifeJourneyList = ljList
	}
	return c.JSON(http.StatusOK, lifeJourneyResponse)
}
