package  controllers

import (
	"smartvitals/src/feautures/esp32/application"
	"github.com/gin-gonic/gin"
)

type DeleteEsp32Controller struct {
	deleteEsp32UseCase *application.DeleteEsp32UseCase
}

func NewDeleteEsp32Controller(deleteEsp32UseCase *application.DeleteEsp32UseCase) *DeleteEsp32Controller {
	return &DeleteEsp32Controller{
		deleteEsp32UseCase: deleteEsp32UseCase,
	}
}


func (controller *DeleteEsp32Controller) Run(c *gin.Context) {
	id := c.Param("id")

	err := controller.deleteEsp32UseCase.Execute(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete esp32"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Esp32 deleted successfully",
	})
}



