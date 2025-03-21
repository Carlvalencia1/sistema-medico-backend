package application

import "smartvitals/src/feautures/cases/domain"

type DeleteMedicalCaseUseCase struct {
	medicalCaseRepository domain.IMedicalCase
}

func NewDeleteMedicalCaseUseCase(medicalCaseRepository domain.IMedicalCase) *DeleteMedicalCaseUseCase {
	return &DeleteMedicalCaseUseCase{medicalCaseRepository: medicalCaseRepository}
}

// Delete a medical case by ID
func (d *DeleteMedicalCaseUseCase) Execute(id int) error { // Cambiado de string a int
	err := d.medicalCaseRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
