package main

import (
	_ "github.com/DTS-STN/benefit-service/docs"
	"github.com/DTS-STN/benefit-service/server"
	"os"
)

// @title Benefit Service
// @version 1.0
// @description This service returns information about benefits

// @host [TBD]
// @BasePath /
func main() {
	//Start the service
	server.Main(os.Args)
}
