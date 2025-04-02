package controllers

import (
	"log"
	"smartvitals/src/feautures/esp32/application"
	"smartvitals/src/feautures/esp32/domain/entities"

	
	"github.com/gin-gonic/gin"
)


type CreateEsp32Controller struct {
	uc *application.SaveEsp32UseCase
}

func NewCreateEsp32Controller(uc *application.SaveEsp32UseCase) *CreateEsp32Controller {
	return &CreateEsp32Controller{
		uc: uc,
	}
}

func (ctr *CreateEsp32Controller) Run(ctx *gin.Context) {
	// Crear una nueva instancia de Esp32
	var esp32 entities.Esp32

	// Vincular el JSON recibido a la estructura
	if err := ctx.ShouldBindJSON(&esp32); err != nil {
		log.Printf("error binding JSON: %v", err)
		ctx.JSON(400, gin.H{"error": "Invalid input format"})
		return
	}

	// Guardar la ESP32 (el UUID se genera en el caso de uso)
	savedEsp32, err := ctr.uc.Run(&esp32)
	if err != nil {
		log.Printf("error saving ESP32: %v", err)
		ctx.JSON(500, gin.H{"error": "Could not save ESP32 device"})
		return
	}

	// Devolver la ESP32 guardada con su nuevo UUID
	ctx.JSON(201, savedEsp32)

}
