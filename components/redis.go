package components

import (
	"github.com/go-redis/redis/v8"
)

func NewRedis(connectionString string) *redis.Client {
	options, err := redis.ParseURL(connectionString)
	if err != nil {
		panic(err)
	}
	options.Username = "" // need to set it to empty string since rediscloud is a dummy username
	rdb := redis.NewClient(options)

	return rdb
}
