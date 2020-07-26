package main

import (
	"fmt"
	"net/http"
)

type Client struct {
	mc      chan []byte
	request *http.Request
}

var clients = make(map[Client]struct{})

func (a *App) UpdateClients() {
	fmt.Println("UpdateClinets triggered")
	for c := range clients {
		params := c.request.URL.Query()
		country, ok := params["country"]
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

func RegisterClient(r *http.Request) Client {
	mc := make(chan []byte)
	c := Client{mc: mc, request: r}
	clients[c] = struct{}{}
	fmt.Printf("Registered client: %s", c)

	return c
}

func RemoveClient(c Client) {
	delete(clients, c)
	fmt.Printf("Deleted client: %s", c)
}
