package news

import (
	"neeews/config"

	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
)

type Api struct {
	Redis  *redis.Client
	Rabbit *amqp.Channel
	Conf   *config.Config
}
