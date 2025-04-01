package controllers

import (
	"smartvitals/src/feautures/users/application"
	"smartvitals/src/feautures/users/domain/entities"
		"github.com/gin-gonic/gin"
)


type CreateUserController struct {
	uc *application.SaveUserUseCase
}

func NewCreateUserController(uc *application.SaveUserUseCase) *CreateUserController {
	return &CreateUserController{
		uc: uc,
	}
}

func (ctr *CreateUserController) Run(ctx *gin.Context) {
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := ctr.uc.Run(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, userResponse)
}
