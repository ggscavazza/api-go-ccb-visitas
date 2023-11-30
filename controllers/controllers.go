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
