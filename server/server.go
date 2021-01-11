package server

import (
	"github.com/DTS-STN/benefit-service/config"
	"github.com/DTS-STN/benefit-service/handlers"
	"github.com/DTS-STN/benefit-service/src/benefits"
	"github.com/DTS-STN/benefit-service/src/db"
	"github.com/DTS-STN/benefit-service/src/lifejourneys"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var echoService *echo.Echo

// Main sets up and launches the echo service
func Main(args []string) {
	// Echo instance
	echoService = echo.New()
	service()
}

// service will setup the echo instance and launch the service
func service() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.Connect(config.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// initialize repo
	benefitRepo := benefits.New(db)
	benefits.Service = benefits.NewService(benefitRepo)
	lifeJourneyRepo := lifejourneys.New(db)
	lifejourneys.Service = lifejourneys.NewService(lifeJourneyRepo)

	echoService.Logger.SetLevel(log.DEBUG)

	// Middleware
	echoService.Use(middleware.Recover())

	// Routes
	echoService.GET("/swagger/*", echoSwagger.WrapHandler)
	echoService.GET("/healthcheck", handlers.HandlerService.HealthCheck)
	echoService.GET("/lifejourneys", handlers.HandlerService.LifeJourneys)
	echoService.GET("/lifejourneys/:id/benefits", handlers.HandlerService.LifeJourneyBenefits)
	echoService.GET("/lifejourneys/:id", handlers.HandlerService.LifeJourneys)
	echoService.GET("/benefits", handlers.HandlerService.Benefits)
	echoService.GET("/benefits/:id", handlers.HandlerService.Benefits)
	// Start server
	echoService.Logger.Fatal(echoService.Start(":" + config.Port))
}
