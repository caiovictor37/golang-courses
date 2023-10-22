package models

import (
	db "golang-courses/alura/3-webapp/database"
	"log"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func SelectProducts() []Product {
	db := db.ConnectIntoDatabase()
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

func CreateProduct(name, description string, price float64, quantity int) {
	db := db.ConnectIntoDatabase()
	insertData, err := db.Prepare("INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	log.Println("Values:", name, description, price, quantity)
	insertData.Exec(name, description, price, quantity)

	defer db.Close()
}
