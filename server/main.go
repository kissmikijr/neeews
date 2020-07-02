package main

import (
	"fmt"
	news "neeews/server/api/news"
	"net/http"
)

func main() {
	fmt.Println("Hello World!!")
	news.NewsHandler()

	http.ListenAndServe(":5000", nil)
}