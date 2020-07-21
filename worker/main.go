package main

import (
	"context"
	"encoding/json"
	"fmt"
	"neeews/components"
	"neeews/config"
	"net/http"

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
	redis := components.NewRedis(conf.RedisConnectionString)
	rabbitChannel := components.NewRabbit(conf)

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

	err = redis.Set(ctx, "hu", r, 0).Err()
	if err != nil {
		panic(err)
	}
	err = rabbitChannel.Publish(
		"",
		conf.RabbitQueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Trigger Client update - go"),
		},
	)
	if err != nil {
		fmt.Printf("Failed to publish message, error: %s", err)
	}
}
