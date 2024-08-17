package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:9000", router))
}
func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "post request received")
}
