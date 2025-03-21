package controllers

import (
	"smartvitals/src/feautures/cases/application"
	"smartvitals/src/feautures/cases/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateMedicalCaseController struct {
	updateMedicalCaseService *application.UpdateMedicalCaseUseCase
}

func NewUpdateMedicalCaseController(updateMedicalCaseService *application.UpdateMedicalCaseUseCase) *UpdateMedicalCaseController {
	return &UpdateMedicalCaseController{updateMedicalCaseService}
}

func (c *UpdateMedicalCaseController) UpdateMedicalCase(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid MedicalCase ID"})
		return
	}

	var medicalCase domain.MedicalCase
	if err := ctx.ShouldBindJSON(&medicalCase); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	medicalCase.IDExpediente = int(id) 
	UpdateMedicalCase, errUpdate := c.updateMedicalCaseService.Execute(medicalCase)
	if errUpdate != nil {
		ctx.JSON(500, gin.H{"error": errUpdate.Error()})
		return
	}

	ctx.JSON(200, UpdateMedicalCase)
}
