package main

import (
	"net/http"

	"github.com/Liedsonfsa/Desafio-Codecon/internal/controllers"
	// "github.com/Liedsonfsa/Desafio-Codecon/internal/models"
)

func main() {
	http.HandleFunc("POST /users", controllers.SaveUsers)

	http.ListenAndServe(":3000", nil)
}