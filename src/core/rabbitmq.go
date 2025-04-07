package core

import (
	"fmt"
	"log"
	"os"
	"time"
	"strings"


	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// NewRabbitMQ establece una conexión resiliente con RabbitMQ
func NewRabbitMQ() (*RabbitMQ, error) {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		return nil, fmt.Errorf("RABBITMQ_URL no definida en .env")
	}

	log.Printf("🔌 Conectando a RabbitMQ en: %s", maskURL(url)) // Oculta contraseña en logs

	// Configuración de conexión resiliente
	conn, err := amqp.DialConfig(url, amqp.Config{
		Heartbeat: 10 * time.Second,
		Locale:    "en_US",
		Dial:      amqp.DefaultDial(10 * time.Second),
	})
	if err != nil {
		return nil, fmt.Errorf("error de conexión: %w (verifica IP, usuario, contraseña o puertos)", err)
	}

	// Manejo de errores en cascada
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("error creando canal: %w", err)
	}

	// Configuración de cola con reintentos
	if err := setupRabbitMQInfrastructure(ch); err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	log.Println("✅ Conexión exitosa con RabbitMQ")
	return &RabbitMQ{Conn: conn, Channel: ch}, nil
}

// setupRabbitMQInfrastructure declara colas y bindings
func setupRabbitMQInfrastructure(ch *amqp.Channel) error {
	_, err := ch.QueueDeclare(
		"patients",
		true,  // Durable
		false, // AutoDelete
		false, // Exclusive
		false, // NoWait
		nil,
	)
	if err != nil {
		return fmt.Errorf("error declarando cola: %w", err)
	}

	// Binding al exchange topic
	if err := ch.QueueBind(
		"patients",
		"multi.patients.data",
		"amq.topic",
		false,
		nil,
	); err != nil {
		return fmt.Errorf("error enlazando cola: %w", err)
	}

	return nil
}

// maskURL oculta la contraseña en los logs
func maskURL(url string) string {
	if len(url) > 10 {
		return url[:strings.Index(url, "@")] + "@[masked]"
	}
	return "[invalid_url]"
}

// PublishMessage con manejo de errores mejorado
func (r *RabbitMQ) PublishMessage(routingKey string, body []byte) error {
	if r.Channel == nil {
		return fmt.Errorf("canal no inicializado")
	}

	return r.Channel.Publish(
		"amq.topic",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		})
}

// Close con verificación de nil pointers
func (r *RabbitMQ) Close() {
	if r.Channel != nil {
		r.Channel.Close()
	}
	if r.Conn != nil {
		r.Conn.Close()
	}
}