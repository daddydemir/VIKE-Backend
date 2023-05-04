package models

import "github.com/google/uuid"

type Verify struct {
	VerifyId   uuid.UUID
	CustomerId uuid.UUID
	Code       string
}
