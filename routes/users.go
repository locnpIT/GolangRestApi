package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"loc.com/hocgolang/models"
)

func signup(context *gin.Context) {
	var user models.User

	//bind json vào User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
