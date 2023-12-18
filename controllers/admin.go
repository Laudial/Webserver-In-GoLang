package controllers

import (
	"net/http"
    "html/template"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("admin/home.html"))
	tpl.Execute(w, nil)
}