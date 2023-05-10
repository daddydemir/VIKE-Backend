package mapper

import (
	"github.com/daddydemir/VIKE-Backend/pkg/models"
	"github.com/google/uuid"
)

type AdvertMapper struct {
	Title        string
	Price        float32
	Stock        int
	NumberOfUser int
	MonthSize    int
}

func (a AdvertMapper) ToAdvert(userId uuid.UUID) models.Advert {
	var advert = models.Advert{
		AdvertId:     uuid.New(),
		CreateUser:   userId,
		Title:        a.Title,
		Price:        a.Price,
		Stock:        a.Stock,
		NumberOfUser: a.NumberOfUser,
		MonthSize:    a.MonthSize,
	}
	return advert
}
