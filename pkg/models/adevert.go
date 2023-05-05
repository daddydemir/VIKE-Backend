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

func (a Advert) Add() Response {
	var r Response
	return r
}

func (a Advert) Update() Response {
	var r Response
	return r
}

func (a Advert) Delete() Response {
	var r Response
	return r
}

func (a Advert) Get() Response {
	var r Response
	return r
}

func (a Advert) List() Response {
	var r Response
	return r
}
