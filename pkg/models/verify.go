package models

import "github.com/google/uuid"

type Verify struct {
	VerifyId   uuid.UUID
	CustomerId uuid.UUID
	Code       string
}

func (v Verify) Add() Response {
	var r Response
	return r
}

func (v Verify) Get() Response {
	var r Response
	return r
}
