package controllers

import (
	"github.com/gin-gonic/gin"
	"smartvitals/src/feautures/users/application"
	"strconv"
)

type GetUserByIdController struct {
	uc *application.GetUserByIDUseCase
}

func NewGetUserByIdController(uc *application.GetUserByIDUseCase) *GetUserByIdController {
	return &GetUserByIdController{uc: uc,}
}

func (ctr *GetUserByIdController) Run(ctx *gin.Context) {
	idStr := ctx.Param("id")
	
	// Convertir el ID a int32
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "ID must be a valid integer",
			"details": err.Error(),
		})
		return
	}

	// Llamar al caso de uso
	user, err := ctr.uc.Run(int32(id))
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "User not found",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})
}