package main

import (
	"os"

	_ "github.com/DTS-STN/benefit-service/docs"
	"github.com/DTS-STN/benefit-service/server"
)

// @title Benefit Service
// @version 1.0
// @description This service returns information about Benefits

// securitydefinitions.oauth2.implicit OAuth2AccessCode
// @in header
// @tokenurl https://keycloak.dev.dts-stn.com/auth/realms/benefit-service-dev/protocol/openid-connect/token
// @authorizationurl https://keycloak.dev.dts-stn.com/auth/realms/benefit-service-dev/protocol/openid-connect/auth

// @host https://benefit-service-dev.dev.dts-stn.com
// @BasePath /
func main() {
	//Start the service
	server.Main(os.Args)
}
