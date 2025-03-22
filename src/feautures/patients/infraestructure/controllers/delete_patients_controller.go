package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"smartvitals/src/feautures/patients/application"
)

type DeletePatientsController struct {
	deletePatientsService *application.DeletePatientsUseCase
}

func NewDeletePatientsController(deletePatientsService *application.DeletePatientsUseCase) *DeletePatientsController {
    return &DeletePatientsController{deletePatientsService}
}

func (c *DeletePatientsController) DeletePatient(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid patient ID"})
		return
	}

	errDelete := c.deletePatientsService.Execute(int(id))
	if errDelete != nil {
		ctx.JSON(500, gin.H{"error": errDelete.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Patient deleted"})
}