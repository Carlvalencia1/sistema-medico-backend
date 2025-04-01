package controllers

import (
	"log"
	"smartvitals/src/feautures/users/application"
	"smartvitals/src/feautures/users/domain/entities"
	"smartvitals/src/core/middlewares"
	"github.com/gin-gonic/gin"
)

type LogInController struct {
	uc *application.LogInUseCase
}

func NewLoginController(uc *application.LogInUseCase) *LogInController {
	return &LogInController{uc: uc}
}

func (ctr *LogInController) Run(ctx *gin.Context) {
	var userLog entities.UserLogIn

	if err := ctx.ShouldBindJSON(&userLog); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	claims, err := ctr.uc.Run(&userLog)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}

	token, errToken := middlewares.GenerateTokenFromClaims(claims)
	if errToken != nil {
		log.Printf("error: %s", errToken)
		ctx.JSON(500, gin.H{"error": "Error generating token"})
		return
	}

	userResponse := entities.UserResponse{
		ID:       claims.ID,
		Name:     claims.Name,
		Rol:      claims.Rol,
		Email:    claims.Email,
		Username: claims.Username,
	}

	ctx.JSON(200, gin.H{
		"token": token,
		"user":  userResponse,
	})
}

