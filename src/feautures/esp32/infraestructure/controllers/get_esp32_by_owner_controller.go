package controllers

import (
	"smartvitals/src/feautures/esp32/application"
	"smartvitals/src/feautures/esp32/domain/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetEsp32ByOwnerController struct {
	uc *application.GetEsp32ByOwnerIDUseCase
}

func NewGetEsp32ByPropietarioController(uc *application.GetEsp32ByOwnerIDUseCase) *GetEsp32ByOwnerController {
	return &GetEsp32ByOwnerController{
		uc: uc,
	}
}

func (c *GetEsp32ByOwnerController) Run(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid id parameter"})
		return
	}

	esp32Devices, err := c.uc.Execute(id) // Cambiado de Run a Execute
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	// Si no hay dispositivos, devolver una lista vacía en lugar de error
	if len(esp32Devices) == 0 {
		ctx.JSON(200, []entities.Esp32{})
		return
	}

	ctx.JSON(200, esp32Devices)
}
