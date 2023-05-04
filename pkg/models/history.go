package models

import (
	"time"

	"github.com/google/uuid"
)

type History struct {
	HistoryId uuid.UUID
	PersonId  uuid.UUID
	Token     string
	Ip        string
	Date      time.Time
}
