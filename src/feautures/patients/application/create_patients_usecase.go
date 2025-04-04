package application

import "smartvitals/src/feautures/patients/domain"

type CreatePatientsUseCase struct {
	patientRepository domain.IPatients
	patientProducer   domain.IPatientsProducer
}

func NewCreatePatientsUseCase(
	patientRepository domain.IPatients,
	patientProducer domain.IPatientsProducer,
) *CreatePatientsUseCase {
	return &CreatePatientsUseCase{
		patientRepository: patientRepository,
		patientProducer:   patientProducer,
	}
}

func (c *CreatePatientsUseCase) Execute(fill domain.Patients) (domain.Patients, error) {
	patientCreated, err := c.patientRepository.Create(fill)
	if err != nil {
		return domain.Patients{}, err
	}

	// Publicar en RabbitMQ (si está configurado)
	_ = c.patientProducer.PublishPatients(patientCreated)

	return patientCreated, nil
}
