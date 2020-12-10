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

	if lifeJourneyRequest.Id == "" {
		lifeJourneyList, err := getAllLifeJourneyBenefits()
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, lifeJourneyResponse)
		}
		lifeJourneyResponse.LifeJourneyList = lifeJourneyList
		return c.JSON(http.StatusOK, lifeJourneyResponse)
	} else {
		lifeJourneyList, err := getLifeJourneyBenefitById(lifeJourneyRequest.Id)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, lifeJourneyResponse)
		}
		lifeJourneyResponse.LifeJourneyList = append(lifeJourneyResponse.LifeJourneyList, lifeJourneyList)
		return c.JSON(http.StatusOK, lifeJourneyResponse.LifeJourneyList)
	}
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

	lifeJourney, err := getLifeJourneyBenefitById(lifeJourneyBenefitsRequest.Id)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, lifeJourneyBenefitsResponse)
	}

	for _, benId := range lifeJourney.RelatedBenefits {
		benefit, err := getBenefitById(benId)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, lifeJourneyBenefitsResponse)
		}
		lifeJourneyBenefitsResponse = append(lifeJourneyBenefitsResponse, benefit)
	}

	return c.JSON(http.StatusOK, lifeJourneyBenefitsResponse)
}

func getLifeJourneyBenefitById(id string) (lifeJourney models.LifeJourney, err error) {

	ljList, err := lifejourneys.LifeJourneyService.LoadLifeJourneys()
	if err != nil {
		return lifeJourney, err
	}

	for _, lj := range ljList {
		if lj.ID == id {
			lifeJourney = lj
			break
		}
	}

	return lifeJourney, nil
}

func getAllLifeJourneyBenefits() (lifeJourneyList []models.LifeJourney, err error) {

	ljList, err := lifejourneys.LifeJourneyService.LoadLifeJourneys()
	if err != nil {
		return lifeJourneyList, err
	}

	return ljList, nil
}

func getBenefitById(benefitId string) (benefit models.Benefits, err error) {

	benList, err := benefits.BenefitsService.LoadBenefits()
	if err != nil {
		return benefit, err
	}

	for _, ben := range benList {
		if ben.ID == benefitId {
			benefit = ben
		}
	}
	return benefit, nil
}
