package main

import (
	"fmt"
	"log"
	"net/http"
//	"github.com/gorilla/mux"
)

var sell_cost float64
var advert_rate float64	
var output =((((sell_cost*0.1)+sell_cost)*(0.1))+(sell_cost * 0.029)+0.3)+(sell_cost * advert_rate)


// some functions
// func calculateFees(W http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")	
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func getVariables() {
	fmt.Println("Enter the sell cost")
	fmt.Scanln(&sell_cost)

	fmt.Println("Enter the advertising rate")
	fmt.Scanln(&advert_rate)
}


func outputVariables() {
	fmt.Printf("%f %f = %f", sell_cost, advert_rate, output)
}


func main() {


// Router and Handler
	//r := mux.NewRouter()
	//r.HandleFunc("/calculate", calculateFees).Methods("GET")

	handleRequests()

//	log.Fatal(http.ListenAndServe(":8000", r))
}