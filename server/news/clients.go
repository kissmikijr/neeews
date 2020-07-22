package news

import (
	"fmt"
	"net/http"
)

type Client struct {
	mc      chan []byte
	request *http.Request
}

var clients = make(map[Client]struct{})

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
		fmt.Println(err)
	}
	forever := make(chan bool)
	go func() {
		for m := range msgs {
			fmt.Printf("Received message: %s", m.Body)

			for c, _ := range clients {
				params := c.request.URL.Query()
				country, ok := params["country"]
				if !ok {
					return
				}

				cNews, err := a.Redis.Get(ctx, country[0]).Result()
				if err != nil {
					fmt.Println(err)
				}
				c.mc <- []byte(cNews)
			}
		}
	}()
	<-forever
}

func RegisterClient(r *http.Request) Client {
	mc := make(chan []byte)
	c := Client{mc: mc, request: r}
	clients[c] = struct{}{}

	return c
}

func RemoveClient(currentClient Client) {
	delete(clients, currentClient)
}
