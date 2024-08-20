package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AmanBhashkar/Banking/domain"
	"github.com/AmanBhashkar/Banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	//wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRespositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:9000", router))
}
func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "post request received")
}
