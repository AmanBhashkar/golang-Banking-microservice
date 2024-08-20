package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/AmanBhashkar/Banking/service"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"pincode" xml:"pincode"`
}
type CustomerHandlers struct {
	service service.CustomerService
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hellow world")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])

}

func (ch *CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "Ashish", City: "New Delhi", Zipcode: "110075"},
	// 	{Name: "Rob", City: "New Delhi", Zipcode: "110076"},
	// }
	customers, _ := ch.service.GetAllCustomer()
	fmt.Print(r.Header.Get("Content-Type"))
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
