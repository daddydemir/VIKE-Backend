package mapper

import (
	"github.com/daddydemir/VIKE-Backend/pkg/encryption"
	"github.com/daddydemir/VIKE-Backend/pkg/models"
	"github.com/google/uuid"
)

type AdminMapper struct {
	Name     string
	Surname  string
	Email    string
	Password string
	Phone    string
}

func (a AdminMapper) ToAdmin() models.Admin {
	var admin = models.Admin{
		Person: models.Person{
			PersonId: uuid.New(),
			Name:     a.Name,
			Surname:  a.Surname,
			Email:    a.Email,
			Password: encryption.HashPass(a.Password),
			Phone:    a.Phone,
			IsActive: true,
		},
	}
	return admin
}
