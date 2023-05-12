package service

import "github.com/daddydemir/VIKE-Backend/pkg/models"

type updateService interface {
	Update() models.Response
}

func UpdateExecute(u updateService) models.Response {
	return u.Update()
}
