package application

import (
	"log"
	"smartvitals/src/feautures/patients/domain"
)

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

	// Intentar publicar en RabbitMQ
	log.Println("Paciente creado, intentando publicar en RabbitMQ...")
	if err := c.patientProducer.PublishPatients(patientCreated); err != nil {
		log.Printf("Error al publicar paciente en RabbitMQ: %v", err)
		// Opcional: puedes decidir si esto es crítico y retornar el error, o continuar igual
	}

	return patientCreated, nil
}
