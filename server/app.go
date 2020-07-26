package main

import (
	"neeews/components"
	"neeews/config"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

type App struct {
	Redis *redis.Client
	Conf  *config.Config
}

func (a *App) InitRouter() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/news/").Subrouter()

	s.HandleFunc("/headlines", a.Headlines).Methods("GET")
	s.HandleFunc("/everything", a.Everything).Methods("GET")
	s.HandleFunc("/countries", a.Countries).Methods("GET")
	s.HandleFunc("/webhook/update-clients", a.HandleUpdateClients).Methods("POST")

	return r
}
func NewApp() *App {
	conf := config.New()
	redis := components.NewRedis(conf.RedisConnectionString)

	return &App{redis, conf}
}
