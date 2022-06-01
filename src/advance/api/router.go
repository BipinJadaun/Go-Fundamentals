package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	router := newRouter()
	http.ListenAndServe(":8080", router)
	fmt.Println("Server started")

}

func newRouter() *mux.Router {
	router := mux.NewRouter()

	// This will map the sprcifild path with the handler method along with the mehod type
	router.HandleFunc("/user", getAll).Methods("GET")
	router.HandleFunc("/user/", getAll).Methods("GET")

	return router
}
