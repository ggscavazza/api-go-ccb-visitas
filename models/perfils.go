package models

type Perfil struct {
	Id     int    `json:"id"`
	Nome   string `json:"nome"`
	Status bool   `json:"status"`
}

var Perfils []Perfil
