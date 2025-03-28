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
