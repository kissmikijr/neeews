package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Body struct {
	Token string
}

var ctx = context.Background()

func (a *App) Headlines(w http.ResponseWriter, r *http.Request) {
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

func (a *App) Everything(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (a *App) HandleUpdateClients(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	secret := strings.Split(token, " ")[1]
	if secret != a.Conf.WorkerToken {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println("Authorized request.")
	a.UpdateClients()
}

func (a *App) Countries(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Data [5]string `json:"data"`
	}{
		a.Conf.CountryCodes,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
