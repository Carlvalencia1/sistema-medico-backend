package application

import "smartvitals/src/feautures/patients/domain"

type DeletePatientsUseCase struct {
	patientRepository domain.IPatients
}

func NewDeletePatientsUseCase(patientRepository domain.IPatients) *DeletePatientsUseCase {
    return &DeletePatientsUseCase{patientRepository: patientRepository}
}

func (d *DeletePatientsUseCase) Execute(id int) error {
	err := d.patientRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}