package controllers

import (
	"fmt"
	"net/http"

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
	fmt.Println("📌 Iniciando GetAllMedicalCaseController...")

	medicalCases, err := c.getAllService.Execute()
	if err != nil {
		fmt.Println("❌ Error en GetAllMedicalCaseController:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting MedicalCases", "details": err.Error()})
		return
	}

	if len(medicalCases) == 0 {
		fmt.Println("⚠️ No se encontraron casos médicos.")
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No medical cases found"})
		return
	}

	fmt.Println("✅ Casos médicos obtenidos con éxito:", medicalCases)
	ctx.JSON(http.StatusOK, medicalCases)
}
