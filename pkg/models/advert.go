package models

import (
	"github.com/daddydemir/VIKE-Backend/config/database"
	"github.com/google/uuid"
	"log"
)

type Advert struct {
	AdvertId     uuid.UUID
	CreateUser   uuid.UUID
	Title        string
	Price        float32
	Stock        int
	NumberOfUser int
	MonthSize    int
	IsActive     bool
}

func (a Advert) Add() Response {
	var r Response
	result := database.Database.Create(&a)
	if result.Error != nil {
		log.Println("an error occurred while saving the advert to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(a)
}

func (a Advert) Update() Response {
	var r Response
	result := database.Database.Save(&a)
	if result.Error != nil {
		log.Println("an error occurred while updating the advert to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(a)
}

func (a Advert) Delete() Response {
	var r Response
	a.IsActive = false
	result := database.Database.Save(&a)
	if result.Error != nil {
		log.Println("an error occurred while deleting the advert to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(a)
}

func (a Advert) Get() Response {
	var r Response
	result := database.Database.Find(&a, "advert_id = ? and is_active = ?", a.AdvertId, true)
	if result.Error != nil {
		log.Println("an error occurred while getting the advert to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(a)
}

func (a Advert) List() Response {
	var r Response
	var adverts []Advert
	result := database.Database.Find(&adverts, "is_active = ?", true)
	if result.Error != nil {
		log.Println("an error occurred while listing the adverts to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(a)
}
