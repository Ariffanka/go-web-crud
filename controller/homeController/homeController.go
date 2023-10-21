package homecontroller

import (
	"html/template"
	"net/http"

)

func Welcome(res http.ResponseWriter, req *http.Request) {
	temp, err:=template.ParseFiles("views/home/index.html")
	if err!= nil{
		panic(err)
	}

	temp.Execute(res, nil)
}
