package handlers

import (
	"github.com/labstack/echo/v4"
)

type HandlerServiceInterface interface {
	HealthCheck(c echo.Context) error
	Benefits(c echo.Context) error
	BenefitsCount(c echo.Context) error
	BenefitsApply(c echo.Context) error
	Questions(c echo.Context) error
  BenefitsEligibility(c echo.Context) error
}

type Handler struct {
}

var HandlerService HandlerServiceInterface = new(Handler)
