package main

import (
	"net/http"

	functions "github.com/kissmikijr/neeews"
)

func main() {
	http.HandleFunc("/headlines", functions.GetHeadlines)
	http.HandleFunc("/everything", functions.GetEverything)
	http.HandleFunc("/countries", functions.GetCountries)
	http.HandleFunc("/scrape-news-api", functions.ScrapeNewsApi)
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":8080", nil)
}
