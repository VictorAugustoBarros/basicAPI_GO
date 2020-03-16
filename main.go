package main

import (
	"basicAPI/controller"
	"basicAPI/model"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := controller.Router
	controller.SetRouters()

	model.PessoaTeste = append(model.PessoaTeste, model.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &model.Address{City: "City X", State: "State X"}})
	model.PessoaTeste = append(model.PessoaTeste, model.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &model.Address{City: "City Z", State: "State Y"}})

	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
