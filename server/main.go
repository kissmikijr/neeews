package main

import (
	"fmt"
	"neeews/components"
	"neeews/config"
	"neeews/server/news"
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
	api.InitRoutes(s)

	go api.UpdateClients()

	fmt.Println("Server listening on port: 5000")
	http.ListenAndServe(":"+conf.Port, r)

}
