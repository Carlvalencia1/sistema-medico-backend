package controllers

import (
    "log"
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

    // Log para verificar que se recibió la solicitud
    log.Println("Iniciando creación de MedicalCase...")

    // Intentar parsear el JSON recibido
    if err := ctx.ShouldBindJSON(&fill); err != nil {
        log.Println("Error al parsear el JSON:", err)
        ctx.JSON(400, gin.H{"error": "Invalid input"})
        return
    }

    // Log para verificar los datos parseados
    log.Printf("Datos recibidos: %+v\n", fill)

    // Intentar ejecutar el caso de uso
    medicalCaseCreate, err := c.createMedicalCaseService.Execute(fill)
    if err != nil {
        log.Println("Error al ejecutar el caso de uso CreateMedicalCase:", err)
        ctx.JSON(500, gin.H{"error": "Error creating MedicalCase"})
        return
    }

    // Log para verificar que el caso médico fue creado exitosamente
    log.Printf("Caso médico creado exitosamente: %+v\n", medicalCaseCreate)

    // Responder con el caso médico creado
    ctx.JSON(201, medicalCaseCreate)
}