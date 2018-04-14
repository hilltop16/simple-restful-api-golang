package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type Person struct {
	// json keys have to be capital letter
	ID string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Address *Address `json:"address"`
}

type Address struct {
	City string `json:"city"`
	State string `json:"state"`

}

var people []Person

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname:"John", Lastname:"Doe", Address:&Address{City:"Chicago", State:"IL"}})
	people = append(people, Person{ID: "2", Firstname:"Ruby", Lastname:"Rails"})
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
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func createPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func deletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}




