package models

import "github.com/google/uuid"

type Person struct {
	PersonId uuid.UUID
	Name     string
	Surname  string
	Email    string
	Password string
	Phone    string
	IsActive bool
}
