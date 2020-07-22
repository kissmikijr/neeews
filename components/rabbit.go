package components

import (
	"neeews/config"

	"github.com/streadway/amqp"
)

func NewRabbit(conf *config.Config) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(conf.RabbitConnectionString)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	_, err = ch.QueueDeclare(
		conf.RabbitQueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	return conn, ch
}
