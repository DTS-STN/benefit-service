package handlers

import (
	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/renderings"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"fmt"
	"encoding/json"
)

type Persons struct {
	Person interface{}
}

// GetQuestions ...
// @Summary Returns question dependencies for requested benefits
// @Description Returns question dependencies for requested benefits
// @ID benefitquestions
// @Success 200 {object} renderings.BenefitQuestionsResponse
// @Router /getquestions [get]
func (h *Handler) GetQuestions(c echo.Context) error {
	getQuestionsRequest := new(bindings.BenefitQuestionsRequest)
	getQuestionsResponse := new(renderings.BenefitQuestionsResponse)

	if err := c.Bind(getQuestionsRequest); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, getQuestionsResponse)
	}

	t := time.Unix(getQuestionsRequest.RequestDate, 0)
	tstring := t.Format("2006-01-02")

	persons := map[string]interface{}{}

	p := map[string]map[string]interface{}{}

	for _, benefit := range getQuestionsRequest.BenefitList {
		p[benefit] = map[string]interface{}{}
		p[benefit][tstring] = nil
	}
	persons["Persons"] = Persons{p}
	benefitJSON, _ := json.Marshal(persons)
	fmt.Println(string(benefitJSON))

	return nil
}
