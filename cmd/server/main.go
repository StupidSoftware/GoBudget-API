package server

import (
	"github.com/breno5g/GoBudget/config"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + PORT)
}
