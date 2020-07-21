package components

import (
	"neeews/config"

	"github.com/streadway/amqp"
)

func NewRabbit(conf *config.Config) *amqp.Channel {
	conn, err := amqp.Dial(conf.RabbitConnectionString)
	if err != nil {
		panic("amqp connection failed")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic("channel connection failed")
	}
	defer ch.Close()
	_, err = ch.QueueDeclare(
		conf.RabbitQueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	return ch
}
