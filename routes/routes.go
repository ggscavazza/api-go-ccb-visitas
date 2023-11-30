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
	r.HandleFunc("/api/perfis", controllers.TodosPerfis).Methods("Get")
	r.HandleFunc("/api/perfis/{id}", controllers.RetornaUmPerfil).Methods("Get")
	r.HandleFunc("/api/perfis", controllers.CriaUmNovoPerfil).Methods("Post")
	r.HandleFunc("/api/perfis/{id}", controllers.DeletaUmPerfil).Methods("Delete")
	r.HandleFunc("/api/perfis/{id}", controllers.EditaPerfil).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
