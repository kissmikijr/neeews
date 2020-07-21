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

func (a *Api) UpdateClients() {
	msgs, err := a.Rabbit.Consume(
		a.Conf.RabbitQueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Error during message consumption")
	}
	forever := make(chan bool)
	go msgHandler(msgs, a.Redis)
	<-forever
}
func msgHandler(msgs <-chan amqp.Delivery, redis *redis.Client) {
	for m := range msgs {
		fmt.Printf("Received message: %s", m.Body)

		for c, _ := range clients {
			params := c.request.URL.Query()
			country, ok := params["country"]
			if !ok {
				return
			}

			cNews, err := redis.Get(ctx, country[0]).Result()
			if err != nil {
				fmt.Println("Error with redis get")
			}
			c.mc <- []byte(cNews)
		}
	}
}
