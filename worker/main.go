package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"neeews/components"
	"neeews/config"
	"net/http"
	"os"
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
	redis := components.NewRedis(conf.RedisConnectionString)

	for _, country := range conf.CountryCodes {

		resp, err := http.Get(fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=%s&apiKey=%s", country, conf.NewsApiKey))
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode >= 400 {
			fmt.Printf("Api call to NewsApi returned with: %d\n", resp.StatusCode)
			var t map[string]interface{}
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			if err := json.Unmarshal(data, &t); err != nil {
				panic(err)
			}
			fmt.Printf("Error message: %s\n", t["message"])
			os.Exit(1)
		}

		var newsApiResponse NewsApiResponse

		err = json.NewDecoder(resp.Body).Decode(&newsApiResponse)
		if err != nil {
			panic(err)
		}
		fmt.Println(newsApiResponse.Articles)
		r, err := json.Marshal(newsApiResponse.Articles)
		if err != nil {
			panic(err)
		}
		err = redis.Set(ctx, country, r, 0).Err()
		if err != nil {
			panic(err)
		}
	}

}
