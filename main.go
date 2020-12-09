package main

import (
	"os"

	_ "github.com/DTS-STN/benefit-service/docs"
	"github.com/DTS-STN/benefit-service/server"
)

// @title Benefit Service
// @version 1.0
// @description This service returns information about Benefits

// @host https://benefit-service-dev.dev.dts-stn.com
// @BasePath /
func main() {
	//Start the service
	server.Main(os.Args)
}
