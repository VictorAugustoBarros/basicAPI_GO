package controller

import (
	"basicAPI/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetPeople -> Retorna todos as pessoas cadastradas
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.PessoaTeste)
}

// GetPerson -> Retorna todos os dados de uma pessoa X
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range model.PessoaTeste {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Person{})
	return
}

// CreatePerson -> Cadastra uma nova pessoa X
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person model.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	model.PessoaTeste = append(model.PessoaTeste, person)
	json.NewEncoder(w).Encode(model.PessoaTeste)
}

// DeletePerson -> Deleta uma pessoa X
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range model.PessoaTeste {
		if item.ID == params["id"] {
			model.PessoaTeste = append(model.PessoaTeste[:index], model.PessoaTeste[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(model.PessoaTeste)
	}
}
