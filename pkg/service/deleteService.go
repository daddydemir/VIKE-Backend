package service

import "github.com/daddydemir/VIKE-Backend/pkg/models"

type deleteService interface {
	Delete() models.Response
}

func DeleteExecute(d deleteService) models.Response {
	return d.Delete()
}
