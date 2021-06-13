package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ScrapeNewsApi(w http.ResponseWriter, r *http.Request) {

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

			w.WriteHeader(resp.StatusCode)
			return
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
		err = redisClient.Set(ctx, country, r, 0).Err()
		if err != nil {
			panic(err)
		}
	}
	w.WriteHeader(http.StatusNoContent)

}
