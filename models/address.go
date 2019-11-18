package models

type Address struct {
	City    string `schema:"cidade" json:"cidade"`
	Cep     string `schema:"cep" json:"cep"`
	ANumber string `schema:"anumber" json:"anumber"`
	AState  string `schema:"estado" json:"estado"`
	Tel     string `schema:"telefone" json:"telefone"`
}
