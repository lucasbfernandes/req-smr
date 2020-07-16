package controllers

import (
	"fmt"
	"req-smr/internal/usecases"

	"github.com/gin-gonic/gin"
)

func SetRequest(context *gin.Context) {
	rawData, err := context.GetRawData()
	if err != nil {
		fmt.Println(err)
		context.JSON(500, gin.H{"error": "Could not get raw data"})
	}

	err = usecases.SetRequest(rawData)
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
