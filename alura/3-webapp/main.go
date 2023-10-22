package main

import (
	"golang-courses/alura/3-webapp/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
