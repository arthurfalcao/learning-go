package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/arthurfalcao/learning-go/web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversão do quantidade:", err)
		}

		models.CreateProduct(name, description, priceConverted, quantityConverted)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	models.DeleteProduct(productID)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	product := models.GetProduct(productID)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço para float:", err)
		}

		idConverted, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do id para int:", err)
		}

		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversão da quantidade para int :", err)
		}

		models.UpdateProduct(idConverted, name, description, priceConverted, quantityConverted)
	}
	http.Redirect(w, r, "/", 301)
}
