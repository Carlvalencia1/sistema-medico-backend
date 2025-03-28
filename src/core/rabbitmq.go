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

	// El exchange amq.topic es un exchange predefinido en RabbitMQ
	// No necesitamos declararlo

	// Declarar la cola
	_, err = ch.QueueDeclare(
		"API2oranges", // nombre de la cola
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("error declarando cola: %w", err)
	}

	// Enlazar la cola al exchange
	err = ch.QueueBind(
		"API2oranges", // nombre de la cola
		"api2.oranges", // routing key
		"amq.topic",   // exchange
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("error enlazando cola: %w", err)
	}

	log.Println("Conexión exitosa con RabbitMQ")
	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}, nil
}

// PublishMessage publica un mensaje en el exchange especificado
func (r *RabbitMQ) PublishMessage(routingKey string, body []byte) error {
	return r.Channel.Publish(
		"amq.topic", // exchange
		routingKey,  // routing key
		false,       // mandatory
		false,       // immediate
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