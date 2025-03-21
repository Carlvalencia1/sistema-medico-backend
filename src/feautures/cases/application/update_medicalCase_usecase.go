package application

import "smartvitals/src/feautures/cases/domain"

type UpdateMedicalCaseUseCase struct {
	medicalCaseRepository domain.IMedicalCase
}

func NewUpdateMedicalCaseUseCase(medicalCaseRepository domain.IMedicalCase) *UpdateMedicalCaseUseCase {
    return &UpdateMedicalCaseUseCase{medicalCaseRepository: medicalCaseRepository}
}

func (u *UpdateMedicalCaseUseCase) Execute(fill domain.MedicalCase) (domain.MedicalCase, error) {
	medicalCaseUpdated, err := u.medicalCaseRepository.Update(fill)
	if err != nil {
		return domain.MedicalCase{}, err
	}
	return medicalCaseUpdated, nil
}