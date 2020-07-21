package main

import (
	"fmt"
	"neeews/components"
	"neeews/config"
	news "neeews/server/api/news"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	conf := config.New()

	rabbitChannel := components.NewRabbit(conf)
	redis := components.NewRedis(conf.RedisConnectionString)

	api := &news.Api{
		Redis:  redis,
		Rabbit: rabbitChannel,
		Conf:   conf,
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/api/news/").Subrouter()
	s.HandleFunc("/headlines", api.Headlines).Methods("GET")
	s.HandleFunc("/everything", api.Everything).Methods("GET")

	go api.UpdateClients()

	fmt.Println("Server listening on port: 5000")
	http.ListenAndServe(":"+conf.Port, r)

}
