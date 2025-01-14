package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/arcanjo96/golang-api-in-memory/lib/db"
	"github.com/arcanjo96/golang-api-in-memory/lib/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao processar JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user.Id = models.SetId()

	db.InMemoryDatabase.Data[user.Id] = user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
