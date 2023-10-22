package controllers

import (
	"golang-courses/alura/3-webapp/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SelectProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// FormValue = attribute "name" of input HTML tag
		name := r.FormValue("name")
		description := r.FormValue("description")
		priceString := r.FormValue("price")
		quantityString := r.FormValue("quantity")

		price, err := strconv.ParseFloat(priceString, 64)
		if err != nil {
			log.Println("Price convertion error:", err)
		}

		quantity, err := strconv.Atoi(quantityString)
		if err != nil {
			log.Println("Quantity convertion error:", err)
		}

		models.CreateProduct(name, description, price, quantity)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
