package handlers

import (
	"github.com/DTS-STN/benefit-service/src/benefits"
	"github.com/DTS-STN/benefit-service/src/lifejourneys"
	"github.com/labstack/echo/v4"
)

type HandlerServiceInterface interface {
	HealthCheck(c echo.Context) error
	LifeJourney(c echo.Context) error
	Benefits(c echo.Context) error
}

type Handler struct {
}

var HandlerService HandlerServiceInterface = new(Handler)

// setLanguage will set the filepath to the correct file for the language requested
func setLanguage(lang string) {
	if lang == "fr" {
		benefits.BenefitsService.SetFilePath("benefit_info_fr.json")
		lifejourneys.LifeJourneyService.SetFilePath("life_journeys_fr.json")
	} else {
		benefits.BenefitsService.SetFilePath("benefit_info_en.json")
		lifejourneys.LifeJourneyService.SetFilePath("life_journeys_en.json")
	}
}
