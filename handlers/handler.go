package handlers

import "github.com/labstack/echo/v4"

// HandlerServiceInterface ...
// Interface for request handlers
type HandlerServiceInterface interface {
	HealthCheck(c echo.Context) error
	GetQuestions(c echo.Context) error
}

// Handler ..
// Generic struct
type Handler struct {
}

// HandlerService ...
// Instance of the HandlerServiceInterface
var HandlerService HandlerServiceInterface = new(Handler)
