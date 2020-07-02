package news

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
)

func NewsHandler() {
	baseUrl := "/news"

	http.HandleFunc(strings.Join(baseUrl, "/headlines"),func (w http.ResponseWriter, r *http.Request){
		b, err := json.Marshal("hurka gyurka")

		if err != nil {
			fmt.Println("Panic.")
		}
		w.Write(b)
	})

	http.HandleFunc("/everything", func(w http.ResponseWriter, r *http.Request){

	})
}