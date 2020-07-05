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
		fmt.Println("amqp connection failed")
	}
	defer conn.Close()

	options, err := redis.ParseURL(conf.RedisConnectionString)
	if err != nil {
		fmt.Println("shit happens")
	}
	options.Username = "" // need to set it to empty string since rediscloud is a dummy username
	rdb := redis.NewClient(options)

	env := &news.Env{
		Redis:  rdb,
		Rabbit: conn,
		Conf:   conf,
	}

	r := mux.NewRouter()
	r.HandleFunc("/headlines", env.Headlines).Methods("GET")
	r.HandleFunc("/everything", env.Everything).Methods("GET")

	go env.UpdateClients()

	fmt.Println("Server listening on port: 5000")
	http.ListenAndServe(":5000", r)

}
