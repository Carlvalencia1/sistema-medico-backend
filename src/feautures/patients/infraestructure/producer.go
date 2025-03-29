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

func (p *Producer) PublishPatients(patients domain.Patients) error {
	// Verificar si rabbitMQ es nil antes de usarlo
	if p.rabbitMQ == nil {
		log.Printf("Advertencia: RabbitMQ no está configurado, saltando publicación")
		return nil // No consideramos esto un error fatal
	}

	jsonData, err := json.Marshal(patients)
	if err != nil {
		log.Printf("Error al convertir patients a JSON: %v", err)
		return err
	}

	err = p.rabbitMQ.PublishMessage("api2.oranges", jsonData)
	if err != nil {
		log.Printf("Error al publicar en RabbitMQ: %v", err)
		return err
	}

	log.Printf("Naranja ID %d publicada exitosamente en RabbitMQ", patients.IDUsuario)
	return nil
}