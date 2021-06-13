package functions

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/kissmikijr/neeews/components"
	"github.com/kissmikijr/neeews/config"
)

var ctx = context.Background()

type Body struct {
	Token string
}
type Source struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type NewsApiResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Cacher interface {
	Get(string) (string, error)
	Set(string, string) error
}
type Cache struct {
	Redis *redis.Client
}

func (c Cache) Get(key string) (string, error) {
	k, err := c.Redis.Get(ctx, key).Result()
	return k, err
}
func (c Cache) Set(key string, value string) error {
	err := c.Redis.Set(ctx, key, value, 0).Err()
	return err
}

var conf *config.Config
var redisClient *redis.Client
var cache Cache

func init() {
	conf = config.New()
	redisClient = components.NewRedis(conf.RedisConnectionString)
	cache = Cache{redisClient}
}
