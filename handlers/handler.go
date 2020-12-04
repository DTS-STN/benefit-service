package handlers

import "github.com/labstack/echo/v4"

type HandlerServiceInterface interface {
	HealthCheck(c echo.Context) error
	LifeJourney(c echo.Context) error
	Benefits(c echo.Context) error
}

type Handler struct {
}

var HandlerService HandlerServiceInterface = new(Handler)
