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

func (g *GetAllMedicalCaseUseCase) Execute()  {

	/*
	cajas, err := g.naranjaRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return cajas, nil
	*/
	}