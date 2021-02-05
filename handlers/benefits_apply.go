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
// @Summary Request redirect url to IEG for a particular benefit
// @ID benefits-apply
// @Accept  json
// @Produce json
// @Security OAuth2AccessCode
// @Header  307 {string} Location "/entity/1"
// @Success 307 {string} string
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Param requestBody body bindings.BenefitApplyRequest true "the benefit you are requesting an apply redirect for"
// @Param Authorization header string true "the bearer token for a particular user"
// @Router /benefits/apply [post]
func (h *Handler) BenefitsApply(c echo.Context) error {
	// check if there is an authorization token
	auth := c.Request().Header.Get("Authorization")
	if auth == "" {
		errObj := new(models.Error)
		errObj.Status = http.StatusForbidden
		errObj.ErrorMessage = "No Auth Token provided"
		return c.JSON(http.StatusForbidden, errObj)
	}

	fmt.Println("something")

	//get the body from the request and try binding it to apply binding
	benefit := new(bindings.BenefitApplyRequest)
	if err := c.Bind(benefit); err != nil {
		errObj := new(models.Error)
		errObj.Status = http.StatusBadRequest
		errObj.ErrorMessage = err.Error()
		return c.JSON(http.StatusBadRequest, errObj)
	}

	if benefit.BenefitType == "" {
		errObj := new(models.Error)
		errObj.Status = http.StatusBadRequest
		errObj.ErrorMessage = "benefitType was not supplied"
		return c.JSON(http.StatusBadRequest, errObj)
	}
	//TODO: take out of handler
	benefitJson, _ := json.Marshal(benefit)
	req, err := http.NewRequest("POST", os.Getenv("CURAM_PRESCREENING_LINK"), bytes.NewBuffer(benefitJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)
	req.Header.Set("guid", "cc6e16b0-db92-459a-91df-f8144befdda9")

	client := interfaces.Client
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("something 2")
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

	if jsonErr := json.Unmarshal(body, &jsonMap); jsonErr != nil {
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

	fmt.Println(iegUrlWithIEGCode)

	return c.Redirect(http.StatusTemporaryRedirect, iegUrlWithIEGCode)
}
