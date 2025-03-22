package application

import "smartvitals/src/feautures/patients/domain"

type CreatePatientsUseCase struct {
	patientRepository domain.IPatients
}

func NewCreatePatientsUseCase(patientRepository domain.IPatients) *CreatePatientsUseCase {
    return &CreatePatientsUseCase{patientRepository: patientRepository}
}

func (c *CreatePatientsUseCase) Execute(fill domain.Patients) (domain.Patients, error) {
	patientCreated, err := c.patientRepository.Create(fill)
	if err != nil {
		return domain.Patients{}, err
	}
	return patientCreated, nil
}