package application

import (
	"smartvitals/src/feautures/cases/domain"
)

type GetAllMedicalCaseUseCase struct {
	medicalCaseRepository domain.IMedicalCase
}

func NewGetAllUseCase(medicalCaseRepository domain.IMedicalCase) *GetAllMedicalCaseUseCase { 
	return &GetAllMedicalCaseUseCase{medicalCaseRepository: medicalCaseRepository} 
}

func (g *GetAllMedicalCaseUseCase) Execute()  ([]domain.MedicalCase, error) {
	medicalCases, err := g.medicalCaseRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return medicalCases, nil
}