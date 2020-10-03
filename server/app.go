package main

import (
	"context"
	"fmt"
	"neeews/components"
	"neeews/config"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var ctx = context.Background()

type Cacher interface {
	Get(string) (string, error)
	Set(string, string) error
}
type Cache struct {
	Redis *redis.Client
}

func (c Cache) Get(key string) (string, error) {
	fmt.Println(key)
	k, err := c.Redis.Get(ctx, key).Result()
	fmt.Println(k)
	return k, err
}
func (c Cache) Set(key string, value string) error {
	err := c.Redis.Set(ctx, key, value, 0).Err()
	return err
}

type App struct {
	Cache Cacher
	Conf  *config.Config
}

func (a *App) InitRouter() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/news/").Subrouter()

	s.HandleFunc("/headlines", a.GetHeadlines).Methods("GET")
	s.HandleFunc("/everything", a.Everything).Methods("GET")
	s.HandleFunc("/countries", a.GetCountries).Methods("GET")

	return r
}
func NewApp() *App {
	conf := config.New()
	redis := components.NewRedis(conf.RedisConnectionString)
	c := Cache{redis}

	return &App{c, conf}
}
