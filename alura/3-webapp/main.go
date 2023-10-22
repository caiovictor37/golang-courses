package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := selectProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func connectIntoDatabase() *sql.DB {
	// Basic implementation, in PRD password should use environment variables
	connection := "user=alura_store dbname=alura_store password=ExamplePsw321. host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func selectProducts() []Product {
	db := connectIntoDatabase()
	selectProducts, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}
	p := Product{}
	products := []Product{}
	for selectProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := selectProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)

	}
	defer db.Close()
	return products
}
