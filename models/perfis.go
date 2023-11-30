package models

type Perfil struct {
	Nome   string `json:"nome"`
	Status bool   `json:"status"`
}

var Perfis []Perfil
