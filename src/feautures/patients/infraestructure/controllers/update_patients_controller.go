package controllers

import (
	"smartvitals/src/feautures/patients/application"
	"smartvitals/src/feautures/patients/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdatePatientsController struct {
	updatePatientsUseCaseService *application.UpdatePatientsUseCase
}

func NewUpdatePatientsController(updatePatientsUseCaseService *application.UpdatePatientsUseCase) *UpdatePatientsController {
    return &UpdatePatientsController{updatePatientsUseCaseService}
}

func (c *UpdatePatientsController) UpdatePatientsUseCase(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid patient ID"})
		return
	}

	var patient domain.Patients
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	patient.IDUsuario = int(id)
	updatePatient, errUpdate := c.updatePatientsUseCaseService.Execute(patient)
	if errUpdate != nil {
		ctx.JSON(500, gin.H{"error": errUpdate.Error()})
		return
	}
	ctx.JSON(200, updatePatient)

}
