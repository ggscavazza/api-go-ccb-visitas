package models

type Ficha struct {
	Id            int    `json:"id"`
	Nome          string `json:"nome"`
	Sobrenome     string `json:"sobrenome"`
	Endereco      string `json:"endereco"`
	Telefone      string `json:"telefone"`
	Whatsapp      string `json:"whatsapp"`
	Observacoes   string `json:"observacoes"`
	Tipo          int    `json:"id_tipo"`
	Ultima_visita string `json:"ultima_visita"`
	Suporte       bool   `json:"suporte"`
	Igreja        int    `json:"id_igreja"`
	Status        bool   `json:"status"`
}

var Fichas []Ficha
