package server

import (
	"github.com/breno5g/GoBudget/config"
	"github.com/breno5g/GoBudget/internal/router"
)

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
