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
	w.WriteHeader(http.StatusCreated)

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
	w.WriteHeader(http.StatusCreated)

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
func TodosPerfils(w http.ResponseWriter, r *http.Request) {
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
	w.WriteHeader(http.StatusCreated)

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

// FUNÇÕES PARA IGREJAS
func TodasIgrejas(w http.ResponseWriter, r *http.Request) {
	var p []models.Igreja
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaIgreja(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var igreja models.Igreja
	database.DB.First(&igreja, id)
	json.NewEncoder(w).Encode(igreja)
}

func CriaUmaNovaIgreja(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	var novaIgreja models.Igreja
	json.NewDecoder(r.Body).Decode(&novaIgreja)
	database.DB.Create(&novaIgreja)
	json.NewEncoder(w).Encode(novaIgreja)
}

func DeletaUmaIgreja(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var igreja models.Igreja
	database.DB.Delete(&igreja, id)
	json.NewEncoder(w).Encode(igreja)
}

func EditaIgreja(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var igreja models.Igreja
	database.DB.First(&igreja, id)
	json.NewDecoder(r.Body).Decode(&igreja)
	database.DB.Save(&igreja)
	json.NewEncoder(w).Encode(igreja)
}

// FUNÇÕES PARA FICHAS
func TodasFichas(w http.ResponseWriter, r *http.Request) {
	var p []models.Ficha
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaFicha(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var ficha models.Ficha
	database.DB.First(&ficha, id)
	json.NewEncoder(w).Encode(ficha)
}

func CriaUmaNovaFicha(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	var novaFicha models.Ficha
	json.NewDecoder(r.Body).Decode(&novaFicha)
	database.DB.Create(&novaFicha)
	json.NewEncoder(w).Encode(novaFicha)
}

func DeletaUmaFicha(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var ficha models.Ficha
	database.DB.Delete(&ficha, id)
	json.NewEncoder(w).Encode(ficha)
}

func EditaFicha(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var ficha models.Ficha
	database.DB.First(&ficha, id)
	json.NewDecoder(r.Body).Decode(&ficha)
	database.DB.Save(&ficha)
	json.NewEncoder(w).Encode(ficha)
}

// FUNÇÕES PARA VISITAS
func TodasVisitas(w http.ResponseWriter, r *http.Request) {
	var p []models.Visita
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaVisita(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var visita models.Visita
	database.DB.First(&visita, id)
	json.NewEncoder(w).Encode(visita)
}

func CriaUmaNovaVisita(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	var novaVisita models.Visita
	json.NewDecoder(r.Body).Decode(&novaVisita)
	database.DB.Create(&novaVisita)
	json.NewEncoder(w).Encode(novaVisita)
}

func DeletaUmaVisita(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var visita models.Visita
	database.DB.Delete(&visita, id)
	json.NewEncoder(w).Encode(visita)
}

func EditaVisita(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var visita models.Visita
	database.DB.First(&visita, id)
	json.NewDecoder(r.Body).Decode(&visita)
	database.DB.Save(&visita)
	json.NewEncoder(w).Encode(visita)
}
