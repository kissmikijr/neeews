package functions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kissmikijr/go-news"
)

func ScrapeNewsApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Scraping started my frieend")
	for _, country := range conf.CountryCodes {
		resp, err := newsApi.TopHeadlines(&news.HeadlinesParameters{Country: country})
		if err != nil {
			panic(err)
		}
		r, err := json.Marshal(resp.Articles)
		if err != nil {
			panic(err)
		}
		err = redisClient.Set(ctx, country, r, 0).Err()
		if err != nil {
			panic(err)
		}
	}
	w.WriteHeader(http.StatusNoContent)

}
