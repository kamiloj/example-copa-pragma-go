package controller

import (
	"awesomeProject/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GrettingService(context *gin.Context){
 	var request model.User
 	if err := context.BindJSON(&request); err != nil{
		return
	}

	var response model.Greet
 	response.MessageGreet = "Welcome "+ request.FirstName + " "+ request.LastName

	context.JSON(http.StatusOK, response)
}