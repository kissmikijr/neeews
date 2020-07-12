package main

import (
	"fmt"
	news "neeews/server/api/news"
	"neeews/server/config"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

func main() {
	conf := config.New()

	conn, err := amqp.Dial(conf.RabbitConnectionString)
	if err != nil {
		panic("amqp connection failed")
	}
	defer conn.Close()

	options, err := redis.ParseURL(conf.RedisConnectionString)
	if err != nil {
		panic("redis connection failed")
	}
	options.Username = "" // need to set it to empty string since rediscloud is a dummy username
	rdb := redis.NewClient(options)

	api := &news.Api{
		Redis:  rdb,
		Rabbit: conn,
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
