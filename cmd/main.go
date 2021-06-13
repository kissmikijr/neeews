package main

import (
	"net/http"

	function "github.com/kissmikijr/neeews"
)

func main() {
	http.HandleFunc("/scrape-news-api", function.ScrapeNewsApi)
	http.ListenAndServe(":8080", nil)
}
