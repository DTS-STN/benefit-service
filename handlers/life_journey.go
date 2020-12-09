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
func (h *Handler) LifeJourney(c echo.Context) error {
	var lifeJourneyResponse = new(renderings.LifeJourneyResponse)
	lifeJourneyRequest := new(bindings.LifeJourneyRequest)

	// Bind the request into our request struct
	if err := c.Bind(lifeJourneyRequest); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, lifeJourneyResponse)
	}

	var lifeJourneyList []models.LifeJourney
	lifeJourneyList, err := getLifeJourneyBenefits(lifeJourneyRequest.Id)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, lifeJourneyResponse)
	}

	lifeJourneyResponse.LifeJourneyList = lifeJourneyList
	return c.JSON(http.StatusOK, lifeJourneyResponse)
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

	if lifeJourneyBenefitsRequest.Id == "" {
		return c.JSON(http.StatusUnprocessableEntity, lifeJourneyBenefitsResponse)
	}

	var lifeJourneyList []models.LifeJourney
	lifeJourneyList, err := getLifeJourneyBenefits(lifeJourneyBenefitsRequest.Id)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, lifeJourneyBenefitsResponse)
	}

	if len(lifeJourneyList) != 1 {
		return c.JSON(http.StatusUnprocessableEntity, lifeJourneyBenefitsResponse)
	}

	lifeJourneyBenefitsResponse, err = getBenefitsByIds(lifeJourneyList[0].RelatedBenefits)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, lifeJourneyBenefitsResponse)
	}

	return c.JSON(http.StatusOK, lifeJourneyBenefitsResponse)
}

func getLifeJourneyBenefits(id string) (lifeJourneyList []models.LifeJourney, err error) {

	ljList, err := lifejourneys.LifeJourneyService.LoadLifeJourneys()
	if err != nil {
		return lifeJourneyList, err
	}

	if id != "" {
		for _, lj := range ljList {
			if lj.ID == id {
				lifeJourneyList = append(lifeJourneyList, lj)
				break
			}

		}
	} else {
		lifeJourneyList = ljList
	}
	return lifeJourneyList, nil
}

func getBenefitsByIds(benefitIds []string) (benefitsList []models.Benefits, err error) {

	benList, err := benefits.BenefitsService.LoadBenefits()
	if err != nil {
		return benefitsList, err
	}

	for _, benId := range benefitIds {
		for _, ben := range benList {
			if ben.ID == benId {
				benefitsList = append(benefitsList, ben)
				break
			}
		}
	}
	return benefitsList, nil
}
