package mapper

import (
	"github.com/daddydemir/VIKE-Backend/pkg/encryption"
	"github.com/daddydemir/VIKE-Backend/pkg/models"
	"github.com/google/uuid"
)

type CustomerMapper struct {
	Name     string
	Surname  string
	Email    string
	Password string
	Phone    string
}

func (c CustomerMapper) ToCustomer() models.Customer {
	var customer = models.Customer{
		Person: models.Person{
			PersonId: uuid.New(),
			Name:     c.Name,
			Surname:  c.Surname,
			Email:    c.Email,
			Password: encryption.HashPass(c.Password),
			Phone:    c.Phone,
			IsActive: true,
		},
	}
	return customer
}
