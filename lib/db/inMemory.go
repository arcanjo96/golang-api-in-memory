package db

import (
	"github.com/arcanjo96/golang-api-in-memory/lib/models"
)

var InMemoryDatabase = &models.Application{
	Data: make(map[models.Id]models.User),
}
