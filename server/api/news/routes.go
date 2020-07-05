package news

import (
	"context"
	"fmt"
	"neeews/server/config"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
)

var ctx = context.Background()

type Env struct {
	Redis  *redis.Client
	Rabbit *amqp.Connection
	Conf   *config.Config
}

func (e *Env) Headlines(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	mc := make(chan []byte)
	currentClient := Client{mc: mc, request: r}
	clients[currentClient] = true

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	country, ok := params["country"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cNews, err := e.Redis.Get(ctx, country[0]).Result()
	if err != nil {
		fmt.Println("Panic.")
	}
	defer func() {
		delete(clients, currentClient)
	}()

	go func() {
		mc <- []byte(cNews)
	}()
	for {
		fmt.Fprintf(w, "data: %s\n\n", <-mc)
	}

}

func (e *Env) Everything(w http.ResponseWriter, r *http.Request) {

}
