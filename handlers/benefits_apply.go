package handlers

import (
	"net/http"

	"github.com/DTS-STN/benefit-service/models"
	"github.com/labstack/echo/v4"
)

// Apply for Benefits
func (h *Handler) BenefitsApply(c echo.Context) error {
	benefit := new(models.Benefits)
	if err := c.Bind(benefit); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, benefit)
}
