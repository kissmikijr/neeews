package main

import (
	"net/http"

	function "github.com/kissmikijr/neeews"
)

func main() {
	http.HandleFunc("/headlines", function.GetHeadlines)
	http.HandleFunc("/everything", function.GetEverything)
	http.HandleFunc("/countries", function.GetCountries)
	http.HandleFunc("/scrape-news-api", function.ScrapeNewsApi)
	http.ListenAndServe(":8080", nil)
}
