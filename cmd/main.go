package main

import (
	"net/http"

	"github.com/Liedsonfsa/Desafio-Codecon/internal/controllers"
)

func main() {
	http.HandleFunc("POST /users", controllers.SaveUsers)
	http.HandleFunc("GET /superusers", controllers.GetSuperUsers)
	http.HandleFunc("GET /top-countries", controllers.GetTopCountries)

	http.ListenAndServe(":3000", nil)
}