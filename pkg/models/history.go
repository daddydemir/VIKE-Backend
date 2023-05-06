package models

import (
	"github.com/daddydemir/VIKE-Backend/config/database"
	"log"
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

func (h History) Add() Response {
	var r Response
	result := database.Database.Create(&h)
	if result.Error != nil {
		log.Println("an error occurred while saving the history to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(h)
}

func (h History) List() Response {
	var r Response
	var histories []History
	result := database.Database.Find(&histories, "person_id = ?", h.PersonId)
	if result.Error != nil {
		log.Println("an error occurred while listing the histories to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(histories)
}
