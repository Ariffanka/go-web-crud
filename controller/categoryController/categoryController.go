package categorycontroller

import (
	"crud-golang-web/entities"
	categorymodel "crud-golang-web/model/categoryModel"
	"html/template"
	"net/http"
	"strconv"
	"time"

)

func Index(res http.ResponseWriter, req *http.Request) {
	categories := categorymodel.GetAll()

	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(res, data)

}

func Add(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		temp, err := template.ParseFiles("views/category/add.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(res, nil)
	}

	if req.Method == "POST" {
		var category entities.Category
		category.Name = req.FormValue("name")
		category.Created_at = time.Now()
		category.Updated_at = time.Now()

		if ok := categorymodel.Create(category); !ok {
			temp, _ := template.ParseFiles("views/category/add")
			temp.Execute(res, nil)
		}

		http.Redirect(res, req, "/categories", http.StatusSeeOther)
	}
}

func Edit(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		temp, err := template.ParseFiles("views/category/edit.html")
		if err != nil {
			panic(err)
		}

		getId := req.URL.Query().Get("id")
		id, err := strconv.Atoi(getId)
		if err != nil {
			panic(err)
		}

		category := categorymodel.Detail(id)
		data := map[string]any{
			"category": category,
		}
		temp.Execute(res, data)
	}

	if req.Method == "POST" {
		var category entities.Category
		getId := req.FormValue("id")
		id, err := strconv.Atoi(getId)
		if err != nil {
			http.Error(res, "Invalid ID", http.StatusBadRequest)
			return
		}

		category.Name = req.FormValue("name")
		category.Updated_at = time.Now()

		if ok := categorymodel.Update(id, category); !ok {
			http.Error(res, "Failed to update category", http.StatusInternalServerError)
			return
		}

		http.Redirect(res, req, "/categories", http.StatusSeeOther)
	}
}

func Delete(res http.ResponseWriter, req *http.Request) {
	getId := req.URL.Query().Get("id")
	id, err := strconv.Atoi(getId)
	if err != nil {
		panic(err)
	}
	if err := categorymodel.Delete(id); err != nil{
		panic(err)
	}

	http.Redirect(res, req, "/categories", http.StatusSeeOther)
}
