package controllers

import (
	"golang-courses/alura/3-webapp/models"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SelectProducts()
	templates.ExecuteTemplate(w, "Index", products)
}
