package main

import (
	"fmt"
	"net/http"
)

func main() {

	app := NewApp()
	r := app.InitRouter()

	fmt.Println("Server listening on port: 5000")
	http.ListenAndServe(":"+app.Conf.Port, r)

}
