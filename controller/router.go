package controller

import "github.com/gorilla/mux"

// Router ->
var Router = mux.NewRouter()

// SetRouters -> seta as rotas padrões
func SetRouters() {
	Router.HandleFunc("/contato", GetPeople).Methods("GET")
	Router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	Router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
	Router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")
}
