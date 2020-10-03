package main

import (
	"fmt"
	"net/http"
)

func main() {

	app := NewApp()
	r := app.InitRouter()

	port := app.Conf.Port

	fmt.Printf("Server listening on port: %s \n", port)
	http.ListenAndServe(":"+port, r)

}
