package routes

import (
	"golang-courses/alura/3-webapp/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
}
