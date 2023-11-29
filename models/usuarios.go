package models

type Usuario struct {
	Email string `json:"login"`
	Senha string `json:"senha"`
}

var Usuarios []Usuario
