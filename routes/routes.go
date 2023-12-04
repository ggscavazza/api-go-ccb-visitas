package routes

import (
	"log"
	"net/http"

	"github.com/ggscavazza/api-go-ccb-visitas/controllers"
	"github.com/ggscavazza/api-go-ccb-visitas/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)

	// Rotas da API para USUARIOS
	r.HandleFunc("/api/usuarios", controllers.TodosUsuarios).Methods("Get")
	r.HandleFunc("/api/usuarios/{id}", controllers.RetornaUmUsuario).Methods("Get")
	r.HandleFunc("/api/usuarios", controllers.CriaUmNovoUsuario).Methods("Post")
	r.HandleFunc("/api/usuarios/{id}", controllers.DeletaUmUsuario).Methods("Delete")
	r.HandleFunc("/api/usuarios/{id}", controllers.EditaUsuario).Methods("Put")

	// Rotas da API para TIPOS
	r.HandleFunc("/api/tipos", controllers.TodosTipos).Methods("Get")
	r.HandleFunc("/api/tipos/{id}", controllers.RetornaUmTipo).Methods("Get")
	r.HandleFunc("/api/tipos", controllers.CriaUmNovoTipo).Methods("Post")
	r.HandleFunc("/api/tipos/{id}", controllers.DeletaUmTipo).Methods("Delete")
	r.HandleFunc("/api/tipos/{id}", controllers.EditaTipo).Methods("Put")

	// Rotas da API para PERFIS
	r.HandleFunc("/api/perfils", controllers.TodosPerfils).Methods("Get")
	r.HandleFunc("/api/perfils/{id}", controllers.RetornaUmPerfil).Methods("Get")
	r.HandleFunc("/api/perfils", controllers.CriaUmNovoPerfil).Methods("Post")
	r.HandleFunc("/api/perfils/{id}", controllers.DeletaUmPerfil).Methods("Delete")
	r.HandleFunc("/api/perfils/{id}", controllers.EditaPerfil).Methods("Put")

	// Rotas da API para IGREJAS
	r.HandleFunc("/api/igrejas", controllers.TodasIgrejas).Methods("Get")
	r.HandleFunc("/api/igrejas/{id}", controllers.RetornaUmaIgreja).Methods("Get")
	r.HandleFunc("/api/igrejas", controllers.CriaUmaNovaIgreja).Methods("Post")
	r.HandleFunc("/api/igrejas/{id}", controllers.DeletaUmaIgreja).Methods("Delete")
	r.HandleFunc("/api/igrejas/{id}", controllers.EditaIgreja).Methods("Put")

	// Rotas da API para FICHAS
	r.HandleFunc("/api/fichas", controllers.TodasFichas).Methods("Get")
	r.HandleFunc("/api/fichas/{id}", controllers.RetornaUmaFicha).Methods("Get")
	r.HandleFunc("/api/fichas", controllers.CriaUmaNovaFicha).Methods("Post")
	r.HandleFunc("/api/fichas/{id}", controllers.DeletaUmaFicha).Methods("Delete")
	r.HandleFunc("/api/fichas/{id}", controllers.EditaFicha).Methods("Put")

	// Rotas da API para VISITAS
	r.HandleFunc("/api/visitas", controllers.TodasVisitas).Methods("Get")
	r.HandleFunc("/api/visitas/{id}", controllers.RetornaUmaVisita).Methods("Get")
	r.HandleFunc("/api/visitas", controllers.CriaUmaNovaVisita).Methods("Post")
	r.HandleFunc("/api/visitas/{id}", controllers.DeletaUmaVisita).Methods("Delete")
	r.HandleFunc("/api/visitas/{id}", controllers.EditaVisita).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
