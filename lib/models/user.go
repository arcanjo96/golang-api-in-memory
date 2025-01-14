package models

import "github.com/google/uuid"

type Id string

type User struct {
	Id        Id     `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Biography string `json:"biography,omitempty"`
}

type Application struct {
	Data map[Id]User
}

func SetId() Id {
	return Id(uuid.New().String())
}
