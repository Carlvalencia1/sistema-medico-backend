package core

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// NewRabbitMQ establece la conexión con RabbitMQ y configura las colas
func NewRabbitMQ() (*RabbitMQ, error) {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		return nil, fmt.Errorf("variable de entorno RABBITMQ_URL no definida")
	}

	// Conectar a RabbitMQ
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("error conectando a RabbitMQ: %w", err)
	}

	// Crear canal
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("error creando canal: %w", err)
	}

	// Declarar la cola para patients
	_, err = ch.QueueDeclare(
		"patients", // Nombre de la cola
		true,       // Durable
		false,      // Auto-delete
		false,      // Exclusive
		false,      // No-wait
		nil,        // Arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("error declarando cola patients: %w", err)
	}

	// Enlazar la cola "patients" al exchange "amq.topic"
	err = ch.QueueBind(
		"patients",            // Nombre de la cola
		"multi.patients.data", // Routing key
		"amq.topic",           // Exchange
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("error enlazando cola patients: %w", err)
	}

	log.Println("Conexión exitosa con RabbitMQ y configuración de cola patients completada")
	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}, nil
}

// PublishMessage publica un mensaje en el exchange especificado
func (r *RabbitMQ) PublishMessage(routingKey string, body []byte) error {
	return r.Channel.Publish(
		"amq.topic", // Exchange
		routingKey,  // Routing key
		false,       // Mandatory
		false,       // Immediate
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		})
}

// Close cierra la conexión y el canal
func (r *RabbitMQ) Close() {
	if r.Channel != nil {
		r.Channel.Close()
	}
	if r.Conn != nil {
		r.Conn.Close()
	}
}
