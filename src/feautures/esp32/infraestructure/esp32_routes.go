package infraestructure

import (
	"smartvitals/src/feautures/esp32/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

type Esp32Routes struct {
	engine                  *gin.Engine
	createEsp32Controller    *controllers.CreateEsp32Controller
	getSp32ByUsernameController *controllers.GetEsp32ByOwnerController
	deleteEsp32Controller *controllers.DeleteEsp32Controller
}

func NewEsp32Routes(
	engine *gin.Engine,
	createEsp32Controller *controllers.CreateEsp32Controller,
	getSp32ByUsernameController  *controllers.GetEsp32ByOwnerController,
	deleteEsp32Controller *controllers.DeleteEsp32Controller,
) *Esp32Routes {
	return &Esp32Routes{
		engine:                  engine,
		createEsp32Controller:    createEsp32Controller,
		getSp32ByUsernameController:   getSp32ByUsernameController,
		deleteEsp32Controller: deleteEsp32Controller,
	}
}

func (routes *Esp32Routes) Run() {
	sp32Routes := routes.engine.Group("/esp32")
	{
		sp32Routes.POST("/", routes.createEsp32Controller.Run)
		sp32Routes.GET("/propietario/:id", routes.getSp32ByUsernameController.Run)
		sp32Routes.DELETE("/:id", routes.deleteEsp32Controller.Run)
	}
}