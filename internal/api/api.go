package api

import (
	"req-smr/internal/controllers"

	"github.com/gin-gonic/gin"
)

func StartAPI() {
	router := gin.Default()

	router.POST("/requests", controllers.SetRequest)
	router.GET("/requests", controllers.GetRequest)

	router.Run(":8080")
}
