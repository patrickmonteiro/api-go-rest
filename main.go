package main

import (
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade {
		{ Nome: "Nome 1", Historia: "História 1"},
		{ Nome: "Nome 2", Historia: "História 2"},
	}
	fmt.Println("Iniciando servidor")
	routes.HandleRequest()
}
