package models

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	LogId   uuid.UUID
	AdminId uuid.UUID
	OrderId uuid.UUID
	Process string
	Time    time.Time
}
