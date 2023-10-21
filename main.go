package main

import (
	"crud-golang-web/config"
	categorycontroller "crud-golang-web/controller/categoryController"
	homecontroller "crud-golang-web/controller/homeController"
	productcontroller "crud-golang-web/controller/productController"
	"log"
	"net/http"
)

func main() {
	config.Connect()

	//homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//category
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	//product
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("server running on port 8080")
	http.ListenAndServe(":8000", nil)
}
