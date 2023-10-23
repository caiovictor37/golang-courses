package models

import (
	db "golang-courses/alura/3-webapp/database"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetProducts() []Product {
	db := db.ConnectIntoDatabase()
	selectProducts, err := db.Query("SELECT * FROM products ORDER BY id")
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
	insertData.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectIntoDatabase()
	deleteData, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}
	deleteData.Exec(id)
	defer db.Close()
}

func GetProduct(id string) Product {
	db := db.ConnectIntoDatabase()
	selectProduct, err := db.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}
	product := Product{}
	for selectProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := selectProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	defer db.Close()
	return product
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectIntoDatabase()
	updateData, err := db.Prepare("UPDATE products SET name=$2, description=$3, price=$4, quantity=$5 WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}
	updateData.Exec(id, name, description, price, quantity)
	defer db.Close()
}
