package server

import (
	"github.com/breno5g/GoBudget/config"
	_ "github.com/breno5g/GoBudget/internal/docs"
	"github.com/breno5g/GoBudget/internal/router"
)

// @title GoBudget API
// @version 1.0
// @description API for GoBudget
// @termsOfService http://swagger.io/terms/
// @contact.name Breno Santos
// @contact.email brenosantos@breno5g.dev
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:3333
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Execute() {
	err := config.Init()
	logger := config.GetLogger("main")
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		panic(err)
	}

	PORT := config.GetEnv().WebServerPort
	logger.Infof("Starting server on port %s", PORT)

	router.Init(PORT)
}
