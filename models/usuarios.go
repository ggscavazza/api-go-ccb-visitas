package models

type Usuario struct {
	Id     int    `json:"id"`
	Nome   string `json:"nome"`
	Email  string `json:"email"`
	Login  string `json:"login"`
	Senha  string `json:"senha"`
	Token  string `json:"token"`
	Perfil int    `json:"id_perfil"`
	Igreja int    `json:"id_igreja"`
	Status bool   `json:"status"`
}

var Usuarios []Usuario
