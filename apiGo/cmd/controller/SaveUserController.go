package controller

import (
	"apiGo/cmd/model"
	"apiGo/cmd/repositories"
	"apiGo/cmd/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserResource struct {
	Error         bool
	Message       string
	ErrorResponse string
}

func Save(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := request.Validate(user)

	if result.Error {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Message})
		return
	}

	success, err := repositories.SaveUser(user)

	if !success {
		ctx.JSON(http.StatusBadRequest, UserResource{
			Error:         true,
			Message:       "problem with the db",
			ErrorResponse: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, UserResource{Error: false, Message: "User saved successfully"})
}
