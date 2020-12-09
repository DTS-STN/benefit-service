package server

import (
	"github.com/DTS-STN/benefit-service/handlers"
	"github.com/DTS-STN/benefit-service/src/benefits"
	"github.com/DTS-STN/benefit-service/src/lifejourneys"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var echoService *echo.Echo

func Main(args []string) {
	setupJsonFilePath()
	// Echo instance
	echoService = echo.New()
	service()
}

func setupJsonFilePath() {
	lifejourneys.LifeJourneyService = lifejourneys.LifeJourneyServiceStruct{Filename: "life_journeys.json"}
	benefits.BenefitsService = benefits.BenefitsServiceStruct{Filename: "benefit_info.json"}
}

func service() {
	echoService.Logger.SetLevel(log.DEBUG)

	// Middleware
	echoService.Use(middleware.Recover())

	// Routes
	echoService.GET("/swagger/*", echoSwagger.WrapHandler)
	echoService.GET("/healthcheck", handlers.HandlerService.HealthCheck)
	echoService.GET("/lifejourneys", handlers.HandlerService.LifeJourney)
	echoService.GET("/lifejourneys/:id/benefits", handlers.HandlerService.LifeJourneyBenefits)
	echoService.GET("/benefits", handlers.HandlerService.Benefits)
	// Start server
	echoService.Logger.Fatal(echoService.Start(":8080"))
}
