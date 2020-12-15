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
	lifejourneys.Service = lifejourneys.ServiceStruct{Filename: "life_journeys_en.json"}
	benefits.Service = benefits.ServiceStruct{Filename: "benefit_info_en.json"}
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
	echoService.GET("/lifejourneys/:id", handlers.HandlerService.LifeJourney)
	echoService.GET("/benefits", handlers.HandlerService.Benefits)
	echoService.GET("/benefits/:id", handlers.HandlerService.Benefits)
	// Start server
	echoService.Logger.Fatal(echoService.Start(":8080"))
}
