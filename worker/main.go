package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"neeews/components"
	"neeews/config"
	"net/http"

	"github.com/streadway/amqp"
	"github.com/tidwall/gjson"
)

var ctx = context.Background()

func main() {
	conf := config.New()
	redis := components.NewRedis(conf.RedisConnectionString)
	conn, ch := components.NewRabbit(conf)
	defer conn.Close()
	defer ch.Close()

	for _, country := range conf.CountryCodes {

		resp, err := http.Get(fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=%s&apiKey=%s", country, conf.NewsApiKey))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)

		articles := gjson.Get(buf.String(), "articles")

		r, err := json.Marshal(articles)
		if err != nil {
			panic(err)
		}
		err = redis.Set(ctx, country, r, 0).Err()
		if err != nil {
			panic(err)
		}
	}

	err := ch.Publish(
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
