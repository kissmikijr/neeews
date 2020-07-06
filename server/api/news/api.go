package news

import (
	"neeews/server/config"

	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
)

type Api struct {
	Redis  *redis.Client
	Rabbit *amqp.Connection
	Conf   *config.Config
}
