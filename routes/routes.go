package routes

import (
	"github/rafaellimasnp/aplicacaoweb/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
