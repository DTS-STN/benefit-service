package server

import (
	"testing"

	"github.com/DTS-STN/benefit-service/handlers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type HandlerServiceMock struct {
	mock.Mock
}

func (m *HandlerServiceMock) HealthCheck(c echo.Context) error {
	args := m.Called()
	return args.Error(1)
}
func (m *HandlerServiceMock) Benefits(c echo.Context) error {
	args := m.Called()
	return args.Error(1)
}
func (m *HandlerServiceMock) BenefitsCount(c echo.Context) error {
	args := m.Called()
	return args.Error(1)
}

func (m *HandlerServiceMock) BenefitsApply(c echo.Context) error {
	args := m.Called()
	return args.Error(1)
}

func (m *HandlerServiceMock) BenefitsEligibility(c echo.Context) error {
  args := m.Called()
	return args.Error(1)
}

func (m *HandlerServiceMock) Questions(c echo.Context) error {
	args := m.Called()
	return args.Error(1)
}

// TODO: This doesn't work, need to setup an http client and call the endpoints to run tests
func TestServer(t *testing.T) {
	e := echo.New()

	// Create a Mock for the interface
	handlerMock := new(HandlerServiceMock)
	// Add a mock call request
	handlerMock.On("HealthCheck", e).
		Return(nil)
	// Set the mock to be used by the code
	handlers.HandlerService = handlers.HandlerServiceInterface(handlerMock)

	// TODO: Add Tests

}
