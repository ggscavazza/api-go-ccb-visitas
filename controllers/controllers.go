package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ggscavazza/api-go-ccb-visitas/database"
	"github.com/ggscavazza/api-go-ccb-visitas/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

// FUNÇÕES PARA USUARIOS
func TodosUsuarios(w http.ResponseWriter, r *http.Request) {
	var p []models.Usuario
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var usuario models.Usuario
	database.DB.First(&usuario, id)
	json.NewEncoder(w).Encode(usuario)
}

func CriaUmNovoUsuario(w http.ResponseWriter, r *http.Request) {
	var novoUsuario models.Usuario
	json.NewDecoder(r.Body).Decode(&novoUsuario)
	database.DB.Create(&novoUsuario)
	json.NewEncoder(w).Encode(novoUsuario)
}

func DeletaUmUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var usuario models.Usuario
	database.DB.Delete(&usuario, id)
	json.NewEncoder(w).Encode(usuario)
}

func EditaUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var usuario models.Usuario
	database.DB.First(&usuario, id)
	json.NewDecoder(r.Body).Decode(&usuario)
	database.DB.Save(&usuario)
	json.NewEncoder(w).Encode(usuario)
}

// FUNÇÕES PARA TIPOS
func TodosTipos(w http.ResponseWriter, r *http.Request) {
	var p []models.Tipo
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmTipo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var tipo models.Tipo
	database.DB.First(&tipo, id)
	json.NewEncoder(w).Encode(tipo)
}

func CriaUmNovoTipo(w http.ResponseWriter, r *http.Request) {
	var novoTipo models.Tipo
	json.NewDecoder(r.Body).Decode(&novoTipo)
	database.DB.Create(&novoTipo)
	json.NewEncoder(w).Encode(novoTipo)
}

func DeletaUmTipo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var tipo models.Tipo
	database.DB.Delete(&tipo, id)
	json.NewEncoder(w).Encode(tipo)
}

func EditaTipo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var tipo models.Tipo
	database.DB.First(&tipo, id)
	json.NewDecoder(r.Body).Decode(&tipo)
	database.DB.Save(&tipo)
	json.NewEncoder(w).Encode(tipo)
}

// FUNÇÕES PARA PERFIS
func TodosPerfis(w http.ResponseWriter, r *http.Request) {
	var p []models.Perfil
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmPerfil(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var perfil models.Perfil
	database.DB.First(&perfil, id)
	json.NewEncoder(w).Encode(perfil)
}

func CriaUmNovoPerfil(w http.ResponseWriter, r *http.Request) {
	var novoPerfil models.Perfil
	json.NewDecoder(r.Body).Decode(&novoPerfil)
	database.DB.Create(&novoPerfil)
	json.NewEncoder(w).Encode(novoPerfil)
}

func DeletaUmPerfil(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var perfil models.Perfil
	database.DB.Delete(&perfil, id)
	json.NewEncoder(w).Encode(perfil)
}

func EditaPerfil(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var perfil models.Perfil
	database.DB.First(&perfil, id)
	json.NewDecoder(r.Body).Decode(&perfil)
	database.DB.Save(&perfil)
	json.NewEncoder(w).Encode(perfil)
}
