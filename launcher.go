package main

import (
	"awesomeProject/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/welcome",controller.GrettingService)
	router.Run(":9080")
}

