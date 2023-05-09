package models

import (
	"github.com/daddydemir/VIKE-Backend/config/database"
	"github.com/google/uuid"
	"log"
	"time"
)

type Session struct {
	SessionId uuid.UUID
	Email     string
	CrDate    time.Time
	ExDate    time.Time
	Token     string
	Ip        string
}

func (s Session) Add() Response {
	var r Response
	result := database.Database.Create(&s)
	if result.Error != nil {
		log.Println("an error occurred while saving the session to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(s)
}

func (s Session) Get() Response {
	var r Response
	result := database.Database.Find(&s, "session_id = ?", s.SessionId)
	if result.Error != nil {
		log.Println("an error occurred while getting the session to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(s)
}
