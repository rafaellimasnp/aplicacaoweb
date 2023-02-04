package main

import (
	"github/rafaellimasnp/aplicacaoweb/routes"
	"net/http"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
