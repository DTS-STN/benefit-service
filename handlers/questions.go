package handlers

import (
	"net/http"

	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/renderings"
	"github.com/DTS-STN/benefit-service/src/questions"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Questions(c echo.Context) error {
	var res = new(renderings.QuestionResponse)
	req := new(bindings.QuestionRequest)
	var err error

	// bind the request into our request struct
	if err = c.Bind(req); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	// if an id is passed in, get the question based on it
	if req.ID != "" {
		if res.Question, err = questions.Service.GetByID(req.Lang, req.ID); err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, res)
		}

		return c.JSON(http.StatusOK, res.Question)
	}

	// otherwire return the list of questions
	if res.QuestionList, err = questions.Service.GetAll(req.Lang); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	return c.JSON(http.StatusOK, res.QuestionList)
}
