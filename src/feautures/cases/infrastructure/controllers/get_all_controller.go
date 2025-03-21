package controllers

import (
	"github.com/gin-gonic/gin"
	"smartvitals/src/feautures/cases/application"
)

type GetAllMedicalCaseController struct {
	getAllService *application.GetAllMedicalCaseUseCase
}

func NewGetAllController(getAllService *application.GetAllMedicalCaseUseCase) *GetAllMedicalCaseController {
	return &GetAllMedicalCaseController{getAllService}
}

func (c *GetAllMedicalCaseController) GetAll(ctx *gin.Context) {
	medicalCases, err := c.getAllService.Execute()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error getting MedicalCases"})
		return
	}
	ctx.JSON(200, medicalCases)
}