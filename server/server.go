package server

import (
	"github.com/DTS-STN/benefit-service/handlers"
	"github.com/DTS-STN/benefit-service/src/lifejourneys"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var echoService *echo.Echo

func Main(args []string) {
	setupLifeJourneyFile()
	// Echo instance
	echoService = echo.New()
	service()
}

func setupLifeJourneyFile() {
	lifejourneys.LifeJourneyService = lifejourneys.LifeJourneyServiceStruct{Filename: "life_journeys.json"}
}

func service() {
	echoService.Logger.SetLevel(log.DEBUG)

	// Middleware
	echoService.Use(middleware.Recover())

	// Routes
	echoService.GET("/swagger/*", echoSwagger.WrapHandler)
	echoService.GET("/healthcheck", handlers.HandlerService.HealthCheck)
	echoService.GET("/lifejourney", handlers.HandlerService.LifeJourney)

	// Start server
	echoService.Logger.Fatal(echoService.Start(":8080"))
}
