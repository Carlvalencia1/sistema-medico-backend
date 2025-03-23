package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"smartvitals/src/core"

	medicalcasesUseCases "smartvitals/src/feautures/cases/application"
	medicalcasesInfrastructure "smartvitals/src/feautures/cases/infrastructure"
	medicalcasesControllers "smartvitals/src/feautures/cases/infrastructure/controllers"

	patientsUseCases "smartvitals/src/feautures/patients/application"
	patientsInfrastructure "smartvitals/src/feautures/patients/infraestructure"
	patientsControllers "smartvitals/src/feautures/patients/infraestructure/controllers"
)

type Dependencies struct {
	engine *gin.Engine
}

func NewDependencies() *Dependencies {
	return &Dependencies{
		engine: gin.Default(),
	}
}

func (d *Dependencies) Run() error {
	database := core.NewDatabase()

	// Configuración de dependencias para casos médicos
	medicalCaseDatabase := medicalcasesInfrastructure.NewMySQL(database.Conn)
	getAllMedicalCaseUseCase := medicalcasesUseCases.NewGetAllUseCase(medicalCaseDatabase)
	getAllMedicalCaseController := medicalcasesControllers.NewGetAllController(getAllMedicalCaseUseCase)
	createMedicalCaseUseCase := medicalcasesUseCases.NewCreateMedicalCaseUseCase(medicalCaseDatabase)
	createMedicalCaseController := medicalcasesControllers.NewCreateMedicalCaseController(createMedicalCaseUseCase)
	updateMedicalCaseUseCase := medicalcasesUseCases.NewUpdateMedicalCaseUseCase(medicalCaseDatabase)
	updateMedicalCaseController := medicalcasesControllers.NewUpdateMedicalCaseController(updateMedicalCaseUseCase)
	deleteMedicalCaseUseCase := medicalcasesUseCases.NewDeleteMedicalCaseUseCase(medicalCaseDatabase)
	deleteMedicalCaseController := medicalcasesControllers.NewDeleteMedicalCaseController(deleteMedicalCaseUseCase)
	medicalCasesRoutes := medicalcasesInfrastructure.NewMedicalCaseRoutes(
		d.engine,
		createMedicalCaseController,
		getAllMedicalCaseController,
		updateMedicalCaseController,
		deleteMedicalCaseController,
	)

	// Configuración de dependencias para pacientes
	patientsDatabase := patientsInfrastructure.NewMySQL(database.Conn)
	getAllPatientsUseCase := patientsUseCases.NewGetAllPatientsUseCase(patientsDatabase)
	getAllPatientsController := patientsControllers.NewGetAllController(getAllPatientsUseCase)
	createPatientsUseCase := patientsUseCases.NewCreatePatientsUseCase(patientsDatabase)
	createPatientsController := patientsControllers.NewCreatePatientsController(createPatientsUseCase)
	updatePatientsUseCase := patientsUseCases.NewUpdatePatientsUseCase(patientsDatabase)
	updatePatientsController := patientsControllers.NewUpdatePatientsController(updatePatientsUseCase)
	deletePatientsUseCase := patientsUseCases.NewDeletePatientsUseCase(patientsDatabase)
	deletePatientsController := patientsControllers.NewDeletePatientsController(deletePatientsUseCase)
	patientsRoutes := patientsInfrastructure.NewPatientsRoutes(
		d.engine,
		createPatientsController,  
		getAllPatientsController,  
		updatePatientsController,
		deletePatientsController,
	)

	// Configuración de CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	d.engine.Use(cors.New(config))

	// Configuración de rutas
	medicalCasesRoutes.SetupRoutes()
	patientsRoutes.SetupRoutes()

	return d.engine.Run(":8080")
}
