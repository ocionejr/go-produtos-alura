package main

import (
	"go-produtos-alura/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas();
	http.ListenAndServe(":8000", nil)
}


