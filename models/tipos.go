package models

type Tipo struct {
	Nome   string `json:"nome"`
	Status bool   `json:"status"`
}

var Tipos []Tipo
