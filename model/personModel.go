package model

// Person -> struct responsável pela criação de usuários
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address -> Struct responsavel por armazaenar o address
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// PessoaTeste -> lista do tipo Person
var PessoaTeste []Person
