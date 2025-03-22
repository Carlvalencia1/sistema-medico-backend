package infraestructure


import (
	"github.com/gin-gonic/gin"
	"smartvitals/src/feautures/patients/infraestructure/controllers"
)

type PatientsRoutes struct {
	engine            *gin.Engine
	createPatientsController *controllers.CreatePatientsController
	getAllPatientsController *controllers.GetAllPatientsController
	updatePatientsController *controllers.UpdatePatientsController
	deletePatientsController *controllers.DeletePatientsController
}

func NewPatientsRoutes(engine *gin.Engine, createPatientsController *controllers.CreatePatientsController, getAllPatientsController *controllers.GetAllPatientsController, updatePatientsController *controllers.UpdatePatientsController, deletePatientsController *controllers.DeletePatientsController) *PatientsRoutes {
	return &PatientsRoutes{
		engine:            engine,
		createPatientsController: createPatientsController,
		getAllPatientsController: getAllPatientsController,
		updatePatientsController: updatePatientsController,
		deletePatientsController: deletePatientsController,
	}
}

func (r *PatientsRoutes) SetupRoutes() {
	patients := r.engine.Group("/patients")
	{
		patients.GET("/", r.getAllPatientsController.GetAll)
		patients.POST("/", r.createPatientsController.CreatePatients)
		patients.PUT("/:id", r.updatePatientsController.UpdatePatient)
		patients.DELETE("/:id", r.deletePatientsController.DeletePatient)
	}
}

func (r *PatientsRoutes) Run() error {
	if err := r.engine.Run(":8080"); err != nil {
        return err
    }
    return nil
}