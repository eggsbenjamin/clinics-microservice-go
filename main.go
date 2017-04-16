package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eggsbenjamin/clinics-microservice-go/constants"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/clinics/postcode/{postcode}", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "{}")
	})

	server := &http.Server{
		Handler: router,
		Addr:    constants.URL,
	}

	log.Fatal(server.ListenAndServe())
}
