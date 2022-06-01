package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	router := NewRouter()

	testServer := httptest.NewServer(router)

	resp, err := http.Get(testServer.URL + "/user")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Error("wrong status code", resp.StatusCode)
	}

	defer resp.Body.Close()

	fmt.Println(resp.Body)

}
