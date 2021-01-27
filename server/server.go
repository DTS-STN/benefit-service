package server

import (
	"github.com/DTS-STN/benefit-service/handlers"
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
	echoService.Logger.SetLevel(log.DEBUG)

	// Middleware
	echoService.Use(middleware.Recover())
	echoService.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// Routes
	echoService.GET("/swagger/*", echoSwagger.WrapHandler)
	echoService.GET("/healthcheck", handlers.HandlerService.HealthCheck)
	echoService.GET("/benefits", handlers.HandlerService.Benefits)
	echoService.GET("/benefits/:id", handlers.HandlerService.Benefits)
	echoService.GET("/benefits/count", handlers.HandlerService.BenefitsCount)
	echoService.POST("/benefits/apply", handlers.HandlerService.BenefitsApply)
	echoService.GET("/questions", handlers.HandlerService.Questions)
	echoService.GET("/questions/:id", handlers.HandlerService.Questions)

	// Start server
	echoService.Logger.Fatal(echoService.Start(":8080"))
}
