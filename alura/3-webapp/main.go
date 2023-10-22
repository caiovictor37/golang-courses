package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := connectIntoDatabase()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"Shirt", "Blue", 29, 10},
		{"Notebook", "Red", 1999, 2},
		{"Phone", "Black", 199, 20},
	}
	templates.ExecuteTemplate(w, "Index", products)
}

func connectIntoDatabase() *sql.DB {
	// Basic implementation, in PRD password should use environment variables
	connection := "user=alura_store dbname=alura_store password=ExamplePsw321. host=localhost sslmode=disabled"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
