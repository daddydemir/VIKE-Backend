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

func (o Order) Add() Response {
	var r Response
	return r
}

func (o Order) Update() Response {
	var r Response
	return r
}

func (o Order) Delete() Response {
	var r Response
	return r
}

func (o Order) Get() Response {
	var r Response
	return r
}

func (o Order) List() Response {
	var r Response
	return r
}
