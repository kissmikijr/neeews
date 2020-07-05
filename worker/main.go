package main

import (
	"context"
	"encoding/json"
	"fmt"
	"neeews/server/config"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
)

var ctx = context.Background()

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

func main() {
	conf := config.New()

	options, err := redis.ParseURL(conf.RedisConnectionString)
	if err != nil {
		fmt.Println("shit happens")
	}
	options.Username = "" // need to set it to empty string since rediscloud is a dummy username
	rdb := redis.NewClient(options)

	conn, err := amqp.Dial(conf.RabbitConnectionString)
	if err != nil {
		fmt.Println("amqp connection failed")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("channel connection failed")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		conf.RabbitQueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	resp, err := http.Get(fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=%s&apiKey=%s", "hu", conf.NewsApiKey))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var nar NewsApiResponse
	err = json.NewDecoder(resp.Body).Decode(&nar)
	if err != nil {
		fmt.Println(err)
	}
	r, err := json.Marshal(nar.Articles)
	if err != nil {
		fmt.Println(err)
	}

	err = rdb.Set(ctx, "hu", r, 0).Err()
	if err != nil {
		panic(err)
	}
	body := "Trigger Client update - go"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		fmt.Printf("Failed to publish message, error: %s", err)
	}
}
