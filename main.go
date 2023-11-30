package main

import (
	"fmt"

	"github.com/ggscavazza/api-go-ccb-visitas/database"
	"github.com/ggscavazza/api-go-ccb-visitas/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
