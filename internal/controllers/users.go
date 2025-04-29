package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Liedsonfsa/Desafio-Codecon/internal/models"
)

func SaveUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.Users
	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, err := os.Create("users.json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&users)
}

func GetSuperUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.Users

	file, err := os.Open("users.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var superusers []models.Users
	for _, user := range users {
		if user.Score >= 900 && user.Active {
			superusers = append(superusers, user)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&superusers)
}