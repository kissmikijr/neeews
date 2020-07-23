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

	conn, channel := components.NewRabbit(conf)
	defer conn.Close()
	defer channel.Close()

	redis := components.NewRedis(conf.RedisConnectionString)

	api := &news.Api{
		Redis:  redis,
		Rabbit: channel,
		Conf:   conf,
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/api/news/").Subrouter()
	api.InitRoutes(s)

	fmt.Println("Server listening on port: 5000")
	http.ListenAndServe(":"+conf.Port, r)

}
