package application

import "smartvitals/src/feautures/patients/domain"

type GetAllPatientsUseCase struct {
	patientRepository domain.IPatients
}

func NewGetAllPatientsUseCase(patientRepository domain.IPatients) *GetAllPatientsUseCase {
	return &GetAllPatientsUseCase{patientRepository: patientRepository}
}

func (g *GetAllPatientsUseCase) Execute() ([]domain.Patients, error) {
	patients, err := g.patientRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return patients, nil
}