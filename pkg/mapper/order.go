package mapper

import (
	"github.com/daddydemir/VIKE-Backend/pkg/models"
	"github.com/google/uuid"
	"time"
)

type OrderMapper struct {
	AdvertId     uuid.UUID
	NumberOfUser int
	Status       string
	ExpireDate   time.Time
}

func (o OrderMapper) ToOrder(userId uuid.UUID) models.Order {
	var order = models.Order{
		OrderId:      uuid.New(),
		AdvertId:     o.AdvertId,
		CustomerId:   userId,
		NumberOfUser: o.NumberOfUser,
		Status:       o.Status,
		Price:        1,
		ExpireDate:   o.ExpireDate,
	}
	return order
	// Order price is problem! TODO:
}
