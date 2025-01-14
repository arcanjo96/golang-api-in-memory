package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/arcanjo96/golang-api-in-memory/lib/db"
	"github.com/arcanjo96/golang-api-in-memory/lib/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parseUuid, err := uuid.Parse(chi.URLParam(r, "id"))
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	id := models.Id(parseUuid.String())
	user := db.InMemoryDatabase.Data[id]
	var updateUser models.User

	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		http.Error(w, "Erro ao processar JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if updateUser.FirstName != "" {
		user.FirstName = updateUser.FirstName
	}

	if updateUser.LastName != "" {
		user.LastName = updateUser.LastName
	}

	if updateUser.Biography != "" {
		user.Biography = updateUser.Biography
	}

	db.InMemoryDatabase.Data[id] = user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
