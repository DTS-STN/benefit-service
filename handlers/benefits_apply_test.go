package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/DTS-STN/benefit-service/bindings"
	"github.com/DTS-STN/benefit-service/interfaces"
	"github.com/DTS-STN/benefit-service/models"
	"github.com/DTS-STN/benefit-service/utils/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BenefitsApplyTestSuite struct {
	suite.Suite
	BenefitJSON       string
	BenefitJSONReader *strings.Reader
	EchoObj           *echo.Echo
}

func (suite *BenefitsApplyTestSuite) SetupTest() {
	os.Setenv("CURAM_PRESCREENING_LINK", "https://curam.com")
	os.Setenv("CURAM_IEG_LINK", "https://curam.com/ieg/{IEGCode}")
	interfaces.Client = &mocks.MockClient{}
	benefit := new(bindings.BenefitApplyRequest)
	benefit.BenefitType = "something"
	benefit_json, _ := json.Marshal(benefit)
	suite.BenefitJSON = string(benefit_json)
	suite.BenefitJSONReader = strings.NewReader(suite.BenefitJSON)
	suite.EchoObj = echo.New()
}

func (suite *BenefitsApplyTestSuite) TestBenefitsApply() {
	// test response data
	json_data := `{"iegExecutionId": "123456"}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json_data)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       r,
		}, nil
	}

	req := httptest.NewRequest(http.MethodPost, "/benefits/apply", suite.BenefitJSONReader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "bearer sometoken")

	// Setup response
	rec := httptest.NewRecorder()

	// Create a new Echo Context
	c := suite.EchoObj.NewContext(req, rec)

	if assert.NoError(suite.T(), HandlerService.BenefitsApply(c)) {
		assert.Equal(suite.T(), http.StatusTemporaryRedirect, rec.Code)
		assert.Equal(suite.T(), "https://curam.com/ieg/123456", rec.Header().Get("Location"))
	}
}

func (suite *BenefitsApplyTestSuite) TestBenefitApplyNoAuth() {
	req := httptest.NewRequest(http.MethodPost, "/benefits/apply", suite.BenefitJSONReader)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := suite.EchoObj.NewContext(req, rec)

	errObj := new(models.Error)
	errObj.Status = http.StatusForbidden
	errObj.ErrorMessage = "No Auth Token provided"

	errObjJson, _ := json.Marshal(errObj)

	if assert.NoError(suite.T(), HandlerService.BenefitsApply(c)) {
		assert.Equal(suite.T(), http.StatusForbidden, rec.Code)
		assert.Equal(suite.T(), string(errObjJson)+"\n", rec.Body.String())
	}
}

func (suite *BenefitsApplyTestSuite) TestBenefitApplyInvalidBody() {
	req := httptest.NewRequest(http.MethodPost, "/benefits/apply", strings.NewReader(`{"some": "json"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "bearer sometoken")

	// Setup response
	rec := httptest.NewRecorder()

	// Create a new Echo Context
	c := suite.EchoObj.NewContext(req, rec)

	errObj := new(models.Error)
	errObj.Status = http.StatusBadRequest
	errObj.ErrorMessage = "benefitType was not supplied"

	errObjJson, _ := json.Marshal(errObj)
	if assert.NoError(suite.T(), HandlerService.BenefitsApply(c)) {
		assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
		assert.Equal(suite.T(), string(errObjJson)+"\n", rec.Body.String())
	}
}

func (suite *BenefitsApplyTestSuite) TestBenefitApplyNonOkayCuramResponse() {
	// test response data
	r := ioutil.NopCloser(strings.NewReader("some error"))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       r,
		}, nil
	}

	req := httptest.NewRequest(http.MethodPost, "/benefits/apply", suite.BenefitJSONReader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "bearer sometoken")

	// Setup response
	rec := httptest.NewRecorder()

	// Create a new Echo Context
	c := suite.EchoObj.NewContext(req, rec)

	errObj := new(models.Error)
	errObj.Status = http.StatusInternalServerError
	errObj.ErrorMessage = "some error"

	errObjJson, _ := json.Marshal(errObj)

	if assert.NoError(suite.T(), HandlerService.BenefitsApply(c)) {
		assert.Equal(suite.T(), http.StatusInternalServerError, rec.Code)
		assert.Equal(suite.T(), string(errObjJson)+"\n", rec.Body.String())
	}
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(BenefitsApplyTestSuite))
}
