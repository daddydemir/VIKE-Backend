package service

import (
	"github.com/daddydemir/VIKE-Backend/config/database"
	"github.com/daddydemir/VIKE-Backend/pkg/models"
	"log"
)

func CreateTables() {

	err := database.Database.AutoMigrate(
		//models.Person{},
		models.Admin{},
		models.Customer{},
		models.Verify{},
		models.History{},
		models.Log{},
		models.Order{},
		models.Session{},
		models.Advert{},
	)

	if err != nil {
		log.Fatal("DATABASE: Table creation has fail:", err)
	}
}
