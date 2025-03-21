package infrastructure

import (
	"github.com/gin-gonic/gin"
	"smartvitals/src/feautures/cases/infrastructure/controllers"
)

type MedicalCaseRoutes struct {
	engine               *gin.Engine
	createMedicalCaseController *controllers.CreateMedicalCaseController
	getAllMedicalCaseController *controllers.GetAllMedicalCaseController
	updateMedicalCaseController *controllers.UpdateMedicalCaseController
	deleteMedicalCaseController *controllers.DeleteMedicalCaseController
}

func NewMedicalCaseRoutes(engine *gin.Engine, createMedicalCaseController *controllers.CreateMedicalCaseController, getAllMedicalCaseController *controllers.GetAllMedicalCaseController, updateMedicalCaseController *controllers.UpdateMedicalCaseController, deleteMedicalCaseController *controllers.DeleteMedicalCaseController ) *MedicalCaseRoutes {
	return &MedicalCaseRoutes{
		engine:               engine,
		createMedicalCaseController: createMedicalCaseController,
		getAllMedicalCaseController: getAllMedicalCaseController,
		updateMedicalCaseController: updateMedicalCaseController,
		deleteMedicalCaseController: deleteMedicalCaseController,
	}
}

func (r *MedicalCaseRoutes) SetupRoutes() {
	medical := r.engine.Group("/medicalcases")
	{
		medical.GET("/", r.getAllMedicalCaseController.GetAll)
		medical.POST("/", r.createMedicalCaseController.Create)
		medical.PUT("/:id", r.updateMedicalCaseController.UpdateMedicalCase)
		medical.DELETE("/:id", r.deleteMedicalCaseController.Delete)
	}
}

func (r *MedicalCaseRoutes) Run() error {
	if err := r.engine.Run(":8080"); err != nil {
		return err
	}
	return nil
}
