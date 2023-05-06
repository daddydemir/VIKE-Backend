package models

import (
	"github.com/daddydemir/VIKE-Backend/config/database"
	"log"
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

func (l Log) Add() Response {
	var r Response
	result := database.Database.Create(&l)
	if result.Error != nil {
		log.Println("an error occurred while saving the log to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(l)
}

func (l Log) List() Response {
	var r Response
	var logs []Log
	result := database.Database.Find(&logs, "admin_id = ?", l.AdminId)
	if result.Error != nil {
		log.Println("an error occurred while listing the logs to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(logs)
}
