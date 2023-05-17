package mapper

import (
	"github.com/daddydemir/VIKE-Backend/config/database"
	"github.com/daddydemir/VIKE-Backend/pkg/models"
	"log"
)

type LoginMapper struct {
	Email    string
	Password string
}

func (l LoginMapper) ToCustomer() models.Customer {
	var customer models.Customer
	result := database.Database.Find(&customer, "email = ? ", l.Email)
	if result.Error != nil {
		log.Println("Error : ", result.Error.Error())
	}
	return customer
}

func (l LoginMapper) ToAdmin() models.Admin {
	var admin models.Admin
	result := database.Database.Find(&admin, "email = ? ", l.Email)
	if result.Error != nil {
		log.Println("Error : ", result.Error.Error())
	}
	return admin
}
