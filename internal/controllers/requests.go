package controllers

import (
	"fmt"
	"req-smr/internal/usecases"

	"github.com/gin-gonic/gin"
)

func SetRequest(context *gin.Context) {
	err := usecases.SetRequest()
	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{"error": "Could not set value"})
	}
	context.JSON(201, nil)
}

func GetRequest(context *gin.Context) {
	counter, err := usecases.GetRequest()
	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{"error": "Could not get value"})
	}
	context.JSON(200, *counter)
}
