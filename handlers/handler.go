package handlers

import (
	"github.com/DTS-STN/benefit-service/src/benefits"
	"github.com/DTS-STN/benefit-service/src/lifejourneys"
	"github.com/labstack/echo/v4"
)

type HandlerServiceInterface interface {
	HealthCheck(c echo.Context) error
	LifeJourney(c echo.Context) error
	LifeJourneyBenefits(c echo.Context) error
	Benefits(c echo.Context) error
	GetLanguage() string
	SetLanguage(lang string)
}

type Handler struct {
	Language string
}

var HandlerService HandlerServiceInterface = new(Handler)

// setLanguage will set the filepath to the correct file for the language requested
func (h *Handler) SetLanguage(lang string) {
	if lang == "fr" {
		benefits.BenefitsService.SetFilePath("benefit_info_fr.json")
		lifejourneys.LifeJourneyService.SetFilePath("life_journeys_fr.json")
	} else {
		benefits.BenefitsService.SetFilePath("benefit_info_en.json")
		lifejourneys.LifeJourneyService.SetFilePath("life_journeys_en.json")
	}

	h.Language = lang
}

func (h *Handler) GetLanguage() string {
	return h.Language
}
