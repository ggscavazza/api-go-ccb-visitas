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
	r.HandleFunc("/api/usuarios", controllers.TodosUsuarios).Methods("Get")
	r.HandleFunc("/api/usuarios/{id}", controllers.RetornaUmUsuario).Methods("Get")
	r.HandleFunc("/api/usuarios", controllers.CriaUmNovoUsuario).Methods("Post")
	r.HandleFunc("/api/usuarios/{id}", controllers.DeletaUmUsuario).Methods("Delete")
	r.HandleFunc("/api/usuarios/{id}", controllers.EditaUsuario).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
