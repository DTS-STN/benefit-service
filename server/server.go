package server

import (
	"github.com/DTS-STN/benefit-service/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var echoService *echo.Echo

// Main function
func Main(args []string) {
	// Echo instance
	echoService = echo.New()
	service()
}

func service() {
	echoService.Logger.SetLevel(log.DEBUG)

	// Middleware
	echoService.Use(middleware.Recover())

	// Routes
	echoService.GET("/swagger/*", echoSwagger.WrapHandler)
	echoService.GET("/healthcheck", handlers.HandlerService.HealthCheck)
	echoService.GET("/getquestions", handlers.HandlerService.GetQuestions)

	// Start server
	echoService.Logger.Fatal(echoService.Start(":8080"))
}
