package database

import (
	"log"

	"github.com/daddydemir/VIKE-Backend/pkg/models"
)

func TableCreate() {
	err := Database.AutoMigrate(
		models.Person{},
		models.Admin{},
		models.Customer{},
		models.Verify{},
		models.History{},
		models.Log{},
		models.Order{},
	)

	if err != nil {
		log.Fatal("DATABASE: Table creation has fail:", err)
	}
}
