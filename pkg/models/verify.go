package models

import (
	"github.com/daddydemir/VIKE-Backend/config/database"
	"github.com/google/uuid"
	"log"
)

type Verify struct {
	VerifyId   uuid.UUID
	CustomerId uuid.UUID
	Code       string
}

func (v Verify) Add() Response {
	var r Response
	result := database.Database.Create(&v)
	if result.Error != nil {
		log.Println("an error occurred while saving the verify to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(v)
}

func (v Verify) Get() Response {
	var r Response
	result := database.Database.Find(&v, "customer_id = ?", v.CustomerId)
	if result.Error != nil {
		log.Println("an error occurred while getting the verify to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(v)
}
