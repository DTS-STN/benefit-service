package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/interfaces"
	"github.com/DTS-STN/benefit-service/models"
	"github.com/labstack/echo/v4"
)

// Apply for Benefits
func (h *Handler) BenefitsApply(c echo.Context) error {
	// check if there is an authorization token
	auth := c.Request().Header.Get("Authorization")
	if auth == "" {
		errObj := new(models.Error)
		errObj.Status = http.StatusForbidden
		errObj.ErrorMessage = "No Auth Token provided"
		return c.JSON(http.StatusForbidden, errObj)
	}

	//get the body from the request and try binding it to apply binding
	benefit := new(bindings.BenefitApplyRequest)
	if err := c.Bind(benefit); err != nil {
		errObj := new(models.Error)
		errObj.Status = http.StatusBadRequest
		errObj.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, errObj)
	}

	benefitJson, _ := json.Marshal(benefit)
	req, err := http.NewRequest("POST", os.Getenv("CURAM_PRESCREENING_LINK"), bytes.NewBuffer(benefitJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	client := interfaces.Client
	resp, err := client.Do(req)

	if err != nil {
		errObj := new(models.Error)
		errObj.Status = http.StatusInternalServerError
		errObj.ErrorMessage = err.Error()
		return c.JSON(http.StatusInternalServerError, errObj)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		errObj := new(models.Error)
		errObj.Status = http.StatusInternalServerError
		errObj.ErrorMessage = string(body)
		return c.JSON(http.StatusInternalServerError, errObj)
	}

	jsonMap := make(map[string]interface{})
	jsonErr := json.Unmarshal(body, &jsonMap)

	if jsonErr != nil {
		errObj := new(models.Error)
		errObj.Status = http.StatusInternalServerError
		errObj.ErrorMessage = err.Error()
		c.Logger().Error(jsonErr)
		return c.JSON(http.StatusInternalServerError, errObj)
	}

	ieg, ok := jsonMap["iegExecutionId"]

	if !ok {
		errObj := new(models.Error)
		errObj.Status = http.StatusInternalServerError
		errObj.ErrorMessage = "No valid IEG ExecutionID returned by curam"
		return c.JSON(http.StatusInternalServerError, errObj)
	}

	iegUrl := os.Getenv("CURAM_IEG_LINK")

	iegUrlWithIEGCode := strings.Replace(iegUrl, "{IEGCode}", fmt.Sprint(ieg), -1)

	return c.Redirect(http.StatusTemporaryRedirect, iegUrlWithIEGCode)
}
