package main

import (
	"errors"
	"fmt"
	"neeews/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CacheMock struct {
	data map[string]string
}

func (cm CacheMock) Get(key string) (string, error) {
	var err error
	v, ok := cm.data[key]
	if !ok {
		err = errors.New("Error in getter")
	}
	return v, err
}
func (cm CacheMock) Set(key string, value string) error {
	cm.data[key] = value
	return nil
}
func TestGetHeadlines(t *testing.T) {
	req, _ := http.NewRequest("GET", "/headlines", nil)

	conf := config.New()
	cm := CacheMock{make(map[string]string)}
	cm.Set("hu", "test-data-value")
	aMock := &App{cm, conf}

	q := req.URL.Query()
	q.Add("country", "hu")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(aMock.GetHeadlines)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code %v", status)
	}

	expected := "jonas"
	fmt.Println(rr.Body.String())

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: %v", rr.Body.String())
	}

}

func TestGetCountries(t *testing.T) {
	req, _ := http.NewRequest("GET", "/countries", nil)

	conf := config.New()
	cm := CacheMock{make(map[string]string)}
	cm.Set("hu", "test-data-value")
	aMock := &App{cm, conf}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(aMock.GetCountries)

	handler.ServeHTTP(rr, req)


}
