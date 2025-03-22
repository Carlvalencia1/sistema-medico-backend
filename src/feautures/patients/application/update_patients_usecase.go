package application

import "smartvitals/src/feautures/patients/domain"

type UpdatePatientsUseCase struct {
	patientRepository domain.IPatients
}

func NewUpdatePatientsUseCase(patientRepository domain.IPatients) *UpdatePatientsUseCase {
    return &UpdatePatientsUseCase{patientRepository: patientRepository}
}

func (u *UpdatePatientsUseCase) Execute(fill domain.Patients) (domain.Patients, error) {
	patientUpdated, err := u.patientRepository.Update(fill)
	if err != nil {
		return domain.Patients{}, err
	}
	return patientUpdated, nil
}