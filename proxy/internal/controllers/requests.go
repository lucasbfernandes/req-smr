package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"req-smr/internal/services"
)

func SetRequest(context *gin.Context) {
	request, err := services.BuildRequestObject(context.Request)
	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{"error": "Could not instantiate request object"})
	}

	err = services.SetRequest(request)
	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{"error": "Could not set request"})
	}
	context.JSON(201, nil)
}

func GetRequests(context *gin.Context) {
	entry, err := services.GetRequests()
	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{"error": "Could not get request"})
	}
	context.JSON(200, entry)
}
