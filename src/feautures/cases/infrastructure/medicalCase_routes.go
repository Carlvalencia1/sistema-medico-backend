package infraestructure

import (
	"github.com/gin-gonic/gin"
	"smartvitals/src/feautures/cases/infrastructure/controllers"
)

type MedicalCaseRoutes struct {
	engine               *gin.Engine
	createMedicalCaseController *controllers.CreateMedicalCaseController
}

func NewMedicalCaseRoutes(engine *gin.Engine, createMedicalCaseController *controllers.CreateMedicalCaseController) *MedicalCaseRoutes {
	return &MedicalCaseRoutes{
		engine:               engine,
		createMedicalCaseController: createMedicalCaseController,
	}
}

func (r *MedicalCaseRoutes) SetupRoutes() {
	medical := r.engine.Group("/medicalcases")
	{
		medical.POST("/", r.createMedicalCaseController.Create)
	}
}

func (r *MedicalCaseRoutes) Run() error {
	if err := r.engine.Run(":8081"); err != nil {
		return err
	}
	return nil
}
