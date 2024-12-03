package controllers

import (
	"net/http"
	"salao/models"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosProdutos := models.RetornaProdutos()

	temp.ExecuteTemplate(w, "Index", todosProdutos)
}
