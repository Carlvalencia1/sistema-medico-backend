package infraestructure

import (
	"encoding/json"
	"log"
	"smartvitals/src/core"
	"smartvitals/src/feautures/patients/domain"
)

type Producer struct {
	rabbitMQ *core.RabbitMQ
}

func NewProducer(rabbitMQ *core.RabbitMQ) *Producer {
	return &Producer{
		rabbitMQ: rabbitMQ,
	}
}

func (p *Producer) PublishPatients(patient domain.Patients) error {
	// Verificar si rabbitMQ es nil antes de usarlo
	if p.rabbitMQ == nil {
		log.Printf("Advertencia: RabbitMQ no está configurado, saltando publicación")
		return nil // No consideramos esto un error fatal
	}

	// Convertir a JSON
	jsonData, err := json.Marshal(patient)
	if err != nil {
		log.Printf("Error al convertir patient a JSON: %v", err)
		return err
	}

	// Publicar usando la routing key correcta para patients
	err = p.rabbitMQ.PublishMessage("multi.patients.data", jsonData)
	if err != nil {
		log.Printf("Error al publicar en RabbitMQ: %v", err)
		return err
	}

	log.Printf("Paciente ID %d publicado exitosamente en RabbitMQ", patient.IDUsuario)
	return nil
}
