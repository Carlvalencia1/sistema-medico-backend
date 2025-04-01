package controllers

import (
	"github.com/gin-gonic/gin"
	"smartvitals/src/feautures/users/application"
)


type GetByUsernameController struct {
	uc *application.GetUserByUsernameUseCase
}

func NewGetByUsernameController(uc *application.GetUserByUsernameUseCase) *GetByUsernameController {
	return &GetByUsernameController{uc: uc,}
}

func (ctr *GetByUsernameController) Run(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := ctr.uc.Run(username)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, user)
}
