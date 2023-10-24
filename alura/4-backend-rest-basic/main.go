package main

import (
	"fmt"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {
	fmt.Println("Starting server...")
	models.Personalities = []models.Personality{
		{Name: "Caio Baeta", Story: "He is a Software Engineer"},
		{Name: "Lucas Ribeiro", Story: "He is an Architect"},
	}
	routes.HandleRequest()
}
