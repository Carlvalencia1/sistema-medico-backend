
package controllers

import (
	"github.com/gin-gonic/gin"
	"smartvitals/src/feautures/users/application"
)

type GetUsersController struct {
	uc *application.GetUsersUseCase
}

func NewGetUsersController(uc *application.GetUsersUseCase) *GetUsersController {
	return &GetUsersController{uc: uc}
}

func (ctr *GetUsersController) Run(ctx *gin.Context) {
	users, err := ctr.uc.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, users)
}