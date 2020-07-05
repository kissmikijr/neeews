package news

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
)

type Client struct {
	mc      chan []byte
	request *http.Request
}

var clients = make(map[Client]bool)

func (e *Env) UpdateClients() {
	ch, err := e.Rabbit.Channel()
	if err != nil {
		fmt.Println("channel connection failed")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		e.Conf.RabbitQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go msgHandler(msgs, e.Redis)
	<-forever
}
func msgHandler(msgs <-chan amqp.Delivery, redis *redis.Client) {
	for d := range msgs {
		fmt.Printf("Received message: %s", d.Body)

		for c, _ := range clients {
			params := c.request.URL.Query()
			country, ok := params["country"]
			if !ok {
				return
			}

			cNews, err := redis.Get(ctx, country[0]).Result()
			if err != nil {
				fmt.Println("Panic.")
			}
			c.mc <- []byte(cNews)
		}
	}
}
