package news

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Body struct {
	token string
}

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
		fmt.Println(err)
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

func (a *Api) HandleUpdateClients(w http.ResponseWriter, r *http.Request) {

	var body Body

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(body, "@@@@@")
	a.UpdateClients()
}
