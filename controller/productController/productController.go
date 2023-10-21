package productcontroller

import (
	"crud-golang-web/entities"
	categorymodel "crud-golang-web/model/categoryModel"
	productmodel "crud-golang-web/model/productModel"
	"net/http"
	"strconv"
	"text/template"
	"time"

)

func Index(res http.ResponseWriter, req *http.Request) {
	products := productmodel.GetAll()

	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(res, data)
}

func Detail(res http.ResponseWriter, req *http.Request) {
	getId := req.URL.Query().Get("id")
	id, err := strconv.Atoi(getId)
	if err != nil{
		panic(err)
	}

	product := productmodel.Detail(id)
	data := map[string]any{
		"product" : product,
	}

	temp, err := template.ParseFiles("views/product/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(res, data)

}

func Add(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		temp, err := template.ParseFiles("views/product/add.html")
		if err != nil {
			panic(err)
		}

		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}

		temp.Execute(res, data)
	}

	if req.Method == "POST" {
		var products entities.Product

		categoryID := req.FormValue("category_id")
		stockStr := req.FormValue("stock")
		if categoryID == "" || stockStr == "" {
			http.Error(res, "Category and stock are required fields", http.StatusBadRequest)
			return
		}

		categoyId, err := strconv.Atoi(categoryID)
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(stockStr)
		if err != nil {
			panic(err)
		}

		products.Name = req.FormValue("name")
		products.Category.Id= uint(categoyId)
		products.Stock= int64(stock)
		products.Desc= req.FormValue("desc")
		products.Created_at = time.Now()
		products.Updated_at = time.Now()

		if ok := productmodel.Create(products); !ok {
			http.Redirect(res,req, req.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(res, req, "/products", http.StatusSeeOther)
	}
}

func Edit(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		temp, err := template.ParseFiles("views/product/edit.html")
		if err != nil {
			panic(err)
		}

		getId := req.URL.Query().Get("id")
		id, err := strconv.Atoi(getId)
		if err != nil {
			panic(err)
		}

		product := productmodel.Detail(id)
		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
			"product" :product, 
		}

		temp.Execute(res, data)
	}

	if req.Method == "POST" {
		var products entities.Product

		getId:= req.FormValue("id")
		id, err := strconv.Atoi(getId)
		if err != nil {
			panic(err)
		}

		categoryID := req.FormValue("category_id")
		stockStr := req.FormValue("stock")
		
		if categoryID == "" || stockStr == "" {
			http.Error(res, "Category and stock are required fields", http.StatusBadRequest)
			return
		}
		
		categoyId, err := strconv.Atoi(categoryID)
		if err != nil {
			panic(err)
		}
	
		stock, err := strconv.Atoi(stockStr)
		if err != nil {
			panic(err)
		}

		products.Name = req.FormValue("name")
		products.Category.Id= uint(categoyId)
		products.Stock= int64(stock)
		products.Desc= req.FormValue("desc")
		products.Updated_at = time.Now()

		if ok := productmodel.Update(id,products); !ok {
			http.Redirect(res,req, req.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(res, req, "/products", http.StatusSeeOther)
	}
}

func Delete(res http.ResponseWriter, req *http.Request) {
	getId := req.URL.Query().Get("id")
	id, err := strconv.Atoi(getId)
	if err != nil {
		panic(err)
	}

	if err := productmodel.Delete(id); err != nil{
		panic(err)
	}

	http.Redirect(res,req, "/products", http.StatusSeeOther)
}
