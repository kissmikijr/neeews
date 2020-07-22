package news

import (
	"context"
	"fmt"
	"net/http"
)

var ctx = context.Background()

func (a *Api) Headlines(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	client := RegisterClient(r)

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	country, ok := params["country"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	headlines, err := a.Redis.Get(ctx, country[0]).Result()
	if err != nil {
		fmt.Println("Panic.")
	}

	defer RemoveClient(client)

	go func() {
		client.mc <- []byte(headlines)
	}()

	for {
		fmt.Fprintf(w, "data: %s\n\n", <-client.mc)
	}

}

func (a *Api) Everything(w http.ResponseWriter, r *http.Request) {

}
