package service

import "github.com/daddydemir/VIKE-Backend/pkg/models"

type getService interface {
	Get() models.Response
}

func GetExecute(g getService) models.Response {
	return g.Get()
}
