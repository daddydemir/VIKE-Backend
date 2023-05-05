package models

import (
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
	return r
}

func (l Log) List() Response {
	var r Response
	return r
}
