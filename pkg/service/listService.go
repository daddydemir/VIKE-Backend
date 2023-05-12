package service

import "github.com/daddydemir/VIKE-Backend/pkg/models"

type listService interface {
	List() models.Response
}

func ListExecute(l listService) models.Response {
	return l.List()
}
