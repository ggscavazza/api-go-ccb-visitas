package models

type Tipo struct {
	Id     int    `json:"id"`
	Nome   string `json:"nome"`
	Status bool   `json:"status"`
}

var Tipos []Tipo
