package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// some functions
func calculateFees(W http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
}


func main() {


// Router and Handler
	r := mux.NewRouter()
	r.HandleFunc("/calculate", calculateFees).Methods("GET")


// 
	
	var sell_cost float64
	var advert_rate float64	

	fmt.Println("Enter the sell cost")
	fmt.Scanln(&sell_cost)

	fmt.Println("Enter the advertising rate")
	fmt.Scanln(&advert_rate)


	var output =((((sell_cost*0.1)+sell_cost)*(0.1))+(sell_cost * 0.029)+0.3)+(sell_cost * advert_rate)

	

	fmt.Printf("%f %f = %f", sell_cost, advert_rate, output)

	log.Fatal(http.ListenAndServe(":8000", r))
}