package controllers

import (
	"net/http"
    "html/template"
)

func PublicHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("public/home.html"))
	tpl.Execute(w, nil)
}