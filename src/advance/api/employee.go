package handler

import (
	"fmt"
	"net/http"
)

func getAll(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Hello World!")
	
}

func getById(resp http.ResponseWriter, req *http.Request, id int){

}
