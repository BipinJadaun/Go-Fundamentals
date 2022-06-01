package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T){
	req, err := http.NewRequest("GET", "", nil)

	if(err != nil){
		fmt.Println("error while creating http request")
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	handlerFunc := http.HandlerFunc(getAll)

	handlerFunc.ServeHTTP(recorder, req)

	if status := recorder.Code ; status != http.StatusOK {
		t.Errorf("handler returned wrong status code %v", status)
	}
	result := recorder.Body
	fmt.Println(result)
}

