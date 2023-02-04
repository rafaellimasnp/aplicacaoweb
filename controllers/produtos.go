package controllers

import (
	"github/rafaellimasnp/aplicacaoweb/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "Index", models.BuscaTodosOsProdutos())

}
