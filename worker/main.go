package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"neeews/config"
	"net/http"
)

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
type Body struct {
	token string
}

var ctx = context.Background()

func main() {
	conf := config.New()
	// redis := components.NewRedis(conf.RedisConnectionString)

	// for _, country := range conf.CountryCodes {

	// 	resp, err := http.Get(fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=%s&apiKey=%s", country, conf.NewsApiKey))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	defer resp.Body.Close()

	// 	var newsApiResponse NewsApiResponse

	// 	err = json.NewDecoder(resp.Body).Decode(&newsApiResponse)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	r, err := json.Marshal(newsApiResponse.Articles)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	err = redis.Set(ctx, country, r, 0).Err()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	b, err := json.Marshal(&Body{conf.WorkerToken})
	if err != nil {
		log.Fatalln(err)
	}

	_, err = http.Post(fmt.Sprintf("%s/api/news/webhook/update-clients", conf.HostName), "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Fatalln(err)
	}
}
