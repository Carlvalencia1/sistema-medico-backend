package controllers

import (
	"smartvitals/src/feautures/cases/application"
	"smartvitals/src/feautures/cases/domain"

	"github.com/gin-gonic/gin"
)

type CreateMedicalCaseController struct {
	createMedicalCaseService *application.CreateMedicalCaseUseCase
}

func NewCreateMedicalCaseController(createMedicalCaseService *application.CreateMedicalCaseUseCase) *CreateMedicalCaseController {
	return &CreateMedicalCaseController{createMedicalCaseService}
}

func (c *CreateMedicalCaseController) Create(ctx *gin.Context) {
	var fill domain.MedicalCase
	if err := ctx.ShouldBindJSON(&fill); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	medicalCaseCreate, err := c.createMedicalCaseService.Execute(fill)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error creating MedicalCase"})
		return
	}
	ctx.JSON(201, medicalCaseCreate)
}
