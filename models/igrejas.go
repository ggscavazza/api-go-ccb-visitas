package models

type Igreja struct {
	Id          int    `json:"id"`
	Nome        string `json:"nome"`
	Cidade      string `json:"cidade"`
	Responsavel string `json:"responsavel"`
	Status      bool   `json:"status"`
}

var Igrejas []Igreja
