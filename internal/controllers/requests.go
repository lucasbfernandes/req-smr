package controllers

import (
	"fmt"
	"req-smr/internal/services"
	"req-smr/internal/usecases"

	"github.com/gin-gonic/gin"
)

func SetRequest(context *gin.Context) {
	request, err := services.BuildRequestObject(context.Request)
	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{"error": "Could not instantiate request object"})
	}

	err = usecases.SetRequest(request)
	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{"error": "Could not set request"})
	}
	context.JSON(201, nil)
}

func GetRequests(context *gin.Context) {
	entry, err := usecases.GetRequests()
	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{"error": "Could not get request"})
	}
	context.JSON(200, entry)
}
