package controllers

import (
	"log"
	"smartvitals/src/feautures/patients/application"
	"smartvitals/src/feautures/patients/domain"
	"github.com/gin-gonic/gin"
)

type CreatePatientsController struct {
	createPatientsUseCase *application.CreatePatientsUseCase
}

func NewCreatePatientsController(createPatientsUseCase *application.CreatePatientsUseCase) *CreatePatientsController {
    return &CreatePatientsController{createPatientsUseCase}
}

func (c *CreatePatientsController) CreatePatients(ctx *gin.Context) {
	var fill domain.Patients
	log.Println("CreatePatientsController")
	if err := ctx.ShouldBindJSON(&fill); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	patientCreated, err := c.createPatientsUseCase.Execute(fill)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, patientCreated)
}