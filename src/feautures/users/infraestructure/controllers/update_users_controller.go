package controllers

import (

	"database/sql"
	"github.com/gin-gonic/gin"
	"smartvitals/src/feautures/users/application"
	"smartvitals/src/feautures/users/domain/entities"
	"smartvitals/src/core/middlewares"
	"strconv"
)



type UpdateUserController struct {
	uc *application.UpdateUserUseCase
}

func NewUpdateUserController(uc *application.UpdateUserUseCase) *UpdateUserController {
	return &UpdateUserController{uc: uc}
}

func (ctr *UpdateUserController) Run(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.ID = int32(id)
	hashPassword, errHashPassword := middlewares.HashPassword(user.Password)
	if errHashPassword != nil {
		ctx.JSON(400, gin.H{"error": errHashPassword.Error()})
		return
	}

	user.Password = hashPassword
	updatedUser, errUpdate := ctr.uc.Run(&user)
	if errUpdate != nil {
		if errUpdate == sql.ErrNoRows {
			ctx.JSON(404, gin.H{"error": "User not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": errUpdate.Error()})
		return
	}

	ctx.JSON(200, updatedUser)
}