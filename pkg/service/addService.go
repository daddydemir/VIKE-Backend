package service

import "github.com/daddydemir/VIKE-Backend/pkg/models"

type addService interface {
	Add() models.Response
}

func AddExecute(a addService) models.Response {
	return a.Add()
}
