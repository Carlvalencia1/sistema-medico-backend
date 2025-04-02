package ports

import "smartvitals/src/feautures/esp32/domain/entities"


type IEsp32 interface {
	Save(esp32 *entities.Esp32) (*entities.Esp32, error)
	GetByPropietario(id int) ([]entities.Esp32, error)
	Delete(id string) error
}