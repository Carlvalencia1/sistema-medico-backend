package application
 
import "smartvitals/src/feautures/esp32/domain/ports"

// DeleteEsp32UseCase is a use case for deleting an ESP32 device by its ID.
type DeleteEsp32UseCase struct {
	db ports.IEsp32
}
 
func NewDeleteEsp32UseCase(db ports.IEsp32) *DeleteEsp32UseCase {
	return &DeleteEsp32UseCase{db: db,}
}

func (uc *DeleteEsp32UseCase) Execute(id string) error {
    return uc.db.Delete(id)
}