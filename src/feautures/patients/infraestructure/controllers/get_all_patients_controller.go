package controllers

import (
	"github.com/gin-gonic/gin"
	"smartvitals/src/feautures/patients/application"
)

type GetAllPatientsController struct {
	getAllService *application.GetAllPatientsUseCase
}

func NewGetAllController(getAllService *application.GetAllPatientsUseCase) *GetAllPatientsController {
    return &GetAllPatientsController{getAllService}
}

func (c *GetAllPatientsController) GetAll(ctx *gin.Context) {
    patients, err  := c.getAllService.Execute()
    if err != nil {
		ctx.JSON(500, gin.H{"error": "Error getting Patients"})
		return
	}
	ctx.JSON(200, patients)
}