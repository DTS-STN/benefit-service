package handlers

import (
	"net/http"

	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/models"
	"github.com/DTS-STN/benefit-service/renderings"
	"github.com/DTS-STN/benefit-service/src/benefits"
	"github.com/DTS-STN/benefit-service/src/lifejourneys"
	"github.com/labstack/echo/v4"
)

// LifeJourney
// @Summary Request Info on Life Journey
// @ID life-journey
// @Accept  json
// @Produce json
// @Success 200 {object} renderings.LifeJourneyResponse
// @Failure 400 {object} renderings.BenefitServiceError
// @Failure 404 {object} renderings.BenefitServiceError
// @Failure 500 {object} renderings.BenefitServiceError
// @Router /lifejourney [get]
func (h *Handler) LifeJourneys(c echo.Context) error {
	var lifeJourneyResponse = new(renderings.LifeJourneyResponse)
	lifeJourneyRequest := new(bindings.LifeJourneyRequest)
	var err error

	// Bind the request into our request struct
	if err = c.Bind(lifeJourneyRequest); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, lifeJourneyResponse)
	}

	// if an ID is passed in, return the specific Life Journey
	if lifeJourneyRequest.Id != "" {
		if lifeJourneyResponse.LifeJourney, err = lifejourneys.Service.GetByID(lifeJourneyRequest.Lang, lifeJourneyRequest.Id); err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, lifeJourneyResponse)
		}
		return c.JSON(http.StatusOK, lifeJourneyResponse.LifeJourney)
	}

	// otherwise return all Life Journeys
	lifeJourneyResponse.LifeJourneyList = lifejourneys.Service.GetAll(lifeJourneyRequest.Lang)
	return c.JSON(http.StatusOK, lifeJourneyResponse.LifeJourneyList)
}

// LifeJourneyBenefits
// @Summary Request Info on Life Journey Related Benefits
// @ID life-journey-benefits
// @Accept  json
// @Produce json
// @Success 200 {object} []models.Benefits
// @Failure 400 {object} renderings.BenefitServiceError
// @Failure 404 {object} renderings.BenefitServiceError
// @Failure 500 {object} renderings.BenefitServiceError
// @Router /lifejourneys/:id/benefits [get]
func (h *Handler) LifeJourneyBenefits(c echo.Context) error {
	var lifeJourneyBenefitsResponse []models.Benefits
	lifeJourneyBenefitsRequest := new(bindings.LifeJourneyBenefitsRequest)

	// Bind the request into our request struct
	if err := c.Bind(lifeJourneyBenefitsRequest); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, lifeJourneyBenefitsResponse)
	}

	lifeJourney, err := lifejourneys.Service.GetByID(lifeJourneyBenefitsRequest.Lang, lifeJourneyBenefitsRequest.Id)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, lifeJourneyBenefitsResponse)
	}

	for _, benId := range lifeJourney.RelatedBenefits {
		benefit, err := benefits.Service.GetByID(lifeJourneyBenefitsRequest.Lang, benId)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, lifeJourneyBenefitsResponse)
		}
		lifeJourneyBenefitsResponse = append(lifeJourneyBenefitsResponse, benefit)
	}

	return c.JSON(http.StatusOK, lifeJourneyBenefitsResponse)
}
