package models

import (
	"github.com/google/uuid"
)

type Advert struct {
	AdvertId     uuid.UUID
	CreateUser   uuid.UUID
	Title        string
	Price        float32
	Stock        int
	NumberOfUser int
	MonthSize    int
}
