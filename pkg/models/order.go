package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderId      uuid.UUID
	AdvertId     uuid.UUID
	CustomerId   uuid.UUID
	NumberOfUser int
	Status       string
	Price        float32
	ExpireDate   time.Time
}
