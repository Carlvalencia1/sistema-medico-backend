package application

import (
    "smartvitals/src/feautures/esp32/domain/entities"
	"smartvitals/src/feautures/esp32/domain/ports"
)
// GetEsp32ByOwnerIDUseCase is a use case for retrieving ESP32 devices by their owner's ID.

type GetEsp32ByOwnerIDUseCase struct {
	db ports.IEsp32
}

func NewGetEsp32ByOwnerIDUseCase(db ports.IEsp32) *GetEsp32ByOwnerIDUseCase {
	return &GetEsp32ByOwnerIDUseCase{db: db,}
}

func (uc *GetEsp32ByOwnerIDUseCase) Execute(id int) ([]entities.Esp32, error) {
	esp32s, err := uc.db.GetByPropietario(id)
	if err != nil {
		return nil, err
	}
	return esp32s, nil
}