package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"smartvitals/src/feautures/cases/application"
)

type DeleteMedicalCaseController struct {
	deleteMedicalCaseService *application.DeleteMedicalCaseUseCase
}

func NewDeleteMedicalCaseController(deleteMedicalCaseService *application.DeleteMedicalCaseUseCase) *DeleteMedicalCaseController {
    return &DeleteMedicalCaseController{deleteMedicalCaseService}
}

func (c *DeleteMedicalCaseController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid caja ID"})
		return
	}

	errDelete := c.deleteMedicalCaseService.Execute(int(id))
	if errDelete != nil {
		ctx.JSON(500, gin.H{"error": errDelete.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "MedicalCase deleted"})
}