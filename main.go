package main

import (
	"net/http"

	"github.com/splorg/golang-web-app/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
