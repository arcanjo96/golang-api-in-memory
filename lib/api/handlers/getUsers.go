package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/arcanjo96/golang-api-in-memory/lib/db"
	"github.com/arcanjo96/golang-api-in-memory/lib/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	usersMap := db.InMemoryDatabase.Data
	users := []models.User{}

	for id, user := range usersMap {
		users = append(users, models.User{
			Id:        id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
