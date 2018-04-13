package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type Person struct {
	id string `json:"id"`
	firstName string `json:"firstName"`
	lastName string `json:"lastName"`
	address *Address `json:"address"`
}

type Address struct {
	city string `json:"city"`
	state string `json:"state"`

}

var people []Person

func main() {
	router := mux.NewRouter()
	people = append(people, Person{id: "1", firstName:"John", lastName:"Doe", address:&Address{city:"Chicago", state:"IL"}})
	people = append(people, Person{id: "2", firstName:"Ruby", lastName:"Rails"})
	router.HandleFunc("/people", getPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", getPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", createPersonEndpoint).Methods("post")
	router.HandleFunc("/people/{id}", deletePersonEndpoint).Methods("delete")
	log.Fatal(http.ListenAndServe(":3000",router))
}

func getPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func getPersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func createPersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func deletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}




