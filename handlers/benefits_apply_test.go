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

	"github.com/DTS-STN/benefit-service/interfaces"
	"github.com/DTS-STN/benefit-service/models"
	"github.com/DTS-STN/benefit-service/utils/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func init() {
	interfaces.Client = &mocks.MockClient{}
}

func setupBenefitsApplyTests() func() {
	os.Setenv("CURAM_PRESCREENING_LINK", "https://curam.com")
	os.Setenv("CURAM_IEG_LINK", "https://curam.com/ieg/{IEGCode}")

	return func() {
	}
}
func TestBenefitsApply(t *testing.T) {
	teardownTests := setupBenefitsApplyTests()
	defer teardownTests()

	// test response data
	json_data := `{"iegExecutionId": "123456"}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json_data)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       r,
		}, nil
	}

	// Setup Echo service
	e := echo.New()
	// Setup http request using httptest
	benefit := new(models.Benefits)
	benefit.BenefitType = "something"
	benefit_json, _ := json.Marshal(benefit)
	benefit_json_string_reader := strings.NewReader(string(benefit_json))
	req := httptest.NewRequest(http.MethodPost, "/benefits/apply", benefit_json_string_reader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "bearer sometoken")

	// Setup response
	rec := httptest.NewRecorder()

	// Create a new Echo Context
	c := e.NewContext(req, rec)

	if assert.NoError(t, HandlerService.BenefitsApply(c)) {
		assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
	}

}
