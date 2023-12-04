package models

type Visita struct {
	Id         int    `json:"id"`
	Ficha      int    `json:"id_ficha"`
	Data       string `json:"data"`
	Horario    string `json:"horario"`
	Data_ordem string `json:"data_ordem"`
	Status     bool   `json:"status"`
}

var Visitas []Visita
