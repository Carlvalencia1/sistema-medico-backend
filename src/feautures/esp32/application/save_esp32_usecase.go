package application

import (
	"smartvitals/src/feautures/esp32/domain/entities"
	"smartvitals/src/feautures/esp32/domain/ports"
	"fmt"
	"log"
	"github.com/google/uuid"
)

type SaveEsp32UseCase struct {
	esp32_repository ports.IEsp32
}

func NewSaveEsp32(esp32_repository ports.IEsp32) *SaveEsp32UseCase {
	return &SaveEsp32UseCase{
		esp32_repository: esp32_repository,
	}
}

func (uc *SaveEsp32UseCase) Run(esp32 *entities.Esp32) (*entities.Esp32, error) {
	fullUUID := uuid.New().String()

	shortID := fmt.Sprintf("ESP-%s", fullUUID[:8])

	esp32.Id = shortID


	savedEsp32, err := uc.esp32_repository.Save(esp32)
	if err != nil {
		log.Printf("error to save esp32: %v", err)
		return &entities.Esp32{}, err
	}

	return savedEsp32, nil
}