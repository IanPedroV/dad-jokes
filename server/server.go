package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	fmt.Println("Running server on port 9000...")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", r))
}
