package handlers

import (
	"github.com/DTS-STN/benefit-service/src/benefits"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBenefits(t *testing.T) {
	// Setup Echo service
	e := echo.New()
	//Load Benefits Data File
	benefits.BenefitsService = benefits.BenefitsServiceStruct{Filename: "../benefit_info.json"}
	// Setup http request using httptest
	req := httptest.NewRequest(http.MethodGet, "/benefits", nil)
	// Create a httptest record
	rec := httptest.NewRecorder()
	// Create a new Echo Context
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, HandlerService.Benefits(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
