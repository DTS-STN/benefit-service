package handlers

import (
	"net/http"

	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/models"
	"github.com/DTS-STN/benefit-service/renderings"
	"github.com/DTS-STN/benefit-service/src/questions"
	"github.com/labstack/echo/v4"
)

// Questions
// @Summary Get a list of questions for pre-screening eligibilty
// @ID questions
// @Accept  json
// @Produce json
// @Success 200 {object} renderings.QuestionResponse
// @Failure 400 {object} renderings.BenefitServiceError
// @Failure 404 {object} renderings.BenefitServiceError
// @Failure 500 {object} renderings.BenefitServiceError
// @Param lang query string false "The language the response should be in. Defaults to English. English and French supported."
// @Router /questions [get]
func (h *Handler) Questions(c echo.Context) error {
	var res = new(renderings.QuestionResponse)
	req := new(bindings.QuestionRequest)
	var err error

	// bind the request into our request struct
	if err = c.Bind(req); err != nil {
		errObj := new(models.Error)
		errObj.Status = http.StatusBadRequest
		errObj.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, errObj)
	}

	// if an id is passed in, get the question based on it
	if req.ID != "" {
		if res.Question, err = questions.Service.GetByID(req.Lang, req.ID); err != nil {
			errObj := new(models.Error)
			errObj.Status = http.StatusBadRequest
			errObj.ErrorMessage = err.Error()
			return c.JSON(http.StatusBadRequest, errObj)
		}

		return c.JSON(http.StatusOK, res.Question)
	}

	// otherwire return the list of questions
	if res.QuestionList, err = questions.Service.GetAll(req.Lang); err != nil {
		errObj := new(models.Error)
		errObj.Status = http.StatusBadRequest
		errObj.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, errObj)
	}

	return c.JSON(http.StatusOK, res.QuestionList)
}
