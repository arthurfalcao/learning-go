package models

import (
	"github.com/arthurfalcao/learning-go/web/db"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := db.ConnectWithDB()

	getAllProducts, err := db.Query("SELECT * FROM products ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for getAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = getAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.ID = id
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
	db := db.ConnectWithDB()

	query, err := db.Prepare("INSERT INTO products (name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	query.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectWithDB()

	query, err := db.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	query.Exec(id)
	defer db.Close()
}

func GetProduct(id string) Product {
	db := db.ConnectWithDB()

	query, err := db.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for query.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := query.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.ID = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	defer db.Close()
	return product
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectWithDB()

	query, err := db.Prepare("UPDATE products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5")
	if err != nil {
		panic(err.Error())
	}

	query.Exec(name, description, price, quantity, id)
	defer db.Close()
}
