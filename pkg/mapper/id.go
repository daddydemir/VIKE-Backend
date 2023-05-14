package mapper

import (
	"github.com/daddydemir/VIKE-Backend/pkg/models"
	"github.com/google/uuid"
	"log"
)

type IdMapper struct {
	Id string
}

func (i IdMapper) Parse() uuid.UUID {
	id, err := uuid.Parse(i.Id)
	if err != nil {
		log.Println("your id isn't uuid format.")
	}
	return id
}

func (i IdMapper) ToAdmin() models.Admin {
	return models.Admin{Person: models.Person{PersonId: i.Parse()}}
}

func (i IdMapper) ToCustomer() models.Customer {
	return models.Customer{Person: models.Person{PersonId: i.Parse()}}
}

func (i IdMapper) ToOrder() models.Order {
	return models.Order{CustomerId: i.Parse()}
}

func (i IdMapper) ToAdvert() models.Advert {
	return models.Advert{AdvertId: i.Parse()}
}

func (i IdMapper) ToHistory() models.History {
	return models.History{PersonId: i.Parse()}
}
