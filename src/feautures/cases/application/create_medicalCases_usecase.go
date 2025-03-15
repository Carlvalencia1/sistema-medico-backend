package application

import "smartvitals/src/feautures/cases/domain"

type CreateMedicalCaseUseCase struct {
	medicalCaseRepository domain.IMedicalCase
}

func NewCreateMedicalCaseUseCase(medicalCaseRepository domain.IMedicalCase) *CreateMedicalCaseUseCase {
	return &CreateMedicalCaseUseCase{medicalCaseRepository: medicalCaseRepository}
}

func (c *CreateMedicalCaseUseCase) Execute(fill domain.MedicalCase) (domain.MedicalCase, error) {
	medicalCreada, err := c.medicalCaseRepository.Create(fill)
	if err != nil {
		return domain.MedicalCase{}, err
	}
	return medicalCreada, nil
}