package main

import (
	"fmt"
	"net/url"

	"github.com/google/uuid"
)

type Client struct {
	mc     chan []byte
	params url.Values
	route  string
	id     uuid.UUID
}

var clients = make(map[uuid.UUID]Client)

func (a *App) UpdateClients() {
	fmt.Println("UpdateClinets triggered")
	for _, c := range clients {
		country, ok := c.params["country"]
		if !ok {
			return
		}

		cNews, err := a.Redis.Get(ctx, country[0]).Result()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Updating client: %s", c)
		c.mc <- []byte(cNews)
	}
}

func RegisterClient(p url.Values, r string) Client {
	mc := make(chan []byte)
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	c := Client{mc: mc, params: p, route: r, id: id}
	clients[id] = c
	fmt.Printf("Registered client: %s", c)

	return c
}

func RemoveClient(c Client) {
	delete(clients, c.id)
	fmt.Printf("Deleted client: %s", c)
}
