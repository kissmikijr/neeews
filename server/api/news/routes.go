package news

import (
	"context"
	"fmt"
	"net/http"
)

var ctx = context.Background()

func (a *Api) Headlines(w http.ResponseWriter, r *http.Request) {
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
	cNews, err := a.Redis.Get(ctx, country[0]).Result()
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

func (a *Api) Everything(w http.ResponseWriter, r *http.Request) {

}
