package functions

import (
	"encoding/json"
	"net/http"
)

func GetHeadlines(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	country, ok := params["country"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cachedHeadlines, err := cache.Get(country[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var headlines []Article
	err = json.Unmarshal([]byte(cachedHeadlines), &headlines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	payload, err := json.Marshal(headlines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(payload)

}

func GetEverything(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func GetCountries(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Data [3]string `json:"data"`
	}{
		conf.CountryCodes,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
