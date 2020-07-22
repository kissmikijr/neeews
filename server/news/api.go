package news

import (
	"neeews/config"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

type Api struct {
	Redis  *redis.Client
	Rabbit *amqp.Channel
	Conf   *config.Config
}

func (a *Api) InitRoutes(r *mux.Router) {
	r.HandleFunc("/headlines", a.Headlines).Methods("GET")
	r.HandleFunc("/everything", a.Everything).Methods("GET")
}
