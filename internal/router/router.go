package router

import (
	"github.com/gin-gonic/gin"
)

func Init(PORT string) {

	r := gin.Default()
	initializeRoutes(r)

	r.Run(":" + PORT)
}
