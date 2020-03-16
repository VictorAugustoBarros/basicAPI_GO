package main

import (
	"basicAPI/controller"
	"basicAPI/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	model.PessoaTeste = append(model.PessoaTeste, model.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &model.Address{City: "City X", State: "State X"}})
	model.PessoaTeste = append(model.PessoaTeste, model.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &model.Address{City: "City Z", State: "State Y"}})

	router.HandleFunc("/contato", controller.GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", controller.GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", controller.DeletePerson).Methods("DELETE")

	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
