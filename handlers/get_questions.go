package handlers

import (
	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/renderings"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"fmt"
)

// TraceRequest struct for sending the list of benefits to OpenFisca
// to get dependencies
type TraceRequest struct {
	Persons struct {
		Person struct {
			Benefit struct {
				RequestDate string
			}
		} `json:"person"`
	} `json:"persons"`
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

	traceRequest := new(TraceRequest)
	benefits := new(Benefits)

	for index, value := range getQuestionsRequest.BenefitList {
		fmt.Println(index, value)
		mapA := map[string]interface{value:{tstring:nil}}
	}

	fmt.Println(benefits)

	return nil
}
