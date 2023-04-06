package main

import (
	"encoding/json"
	"fmt"
	"microservice/models"
	"net/http"

	"github.com/go-chi/chi"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var users []models.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/users/%s", id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var user models.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
