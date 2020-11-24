package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetQuestions ...
// @Summary Returns question dependencies for requested benefits
// @Description Returns question dependencies for requested benefits
// @ID benefitquestions
// @Success 200 {object} renderings.BenefitQuestionsResponse
// @Router /getquestions [get]
func (h *Handler) GetQuestions(c echo.Context) error {
	return c.String(http.StatusOK, "Healthy")
}
