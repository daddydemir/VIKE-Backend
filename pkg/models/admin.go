package models

import (
	"github.com/daddydemir/VIKE-Backend/config/database"
	"log"
)

type Admin struct {
	Person
}

func (a Admin) Add() Response {
	var r Response
	result := database.Database.Create(&a)
	if result.Error != nil {
		log.Println("an error occurred while saving the admin to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(a)
}

func (a Admin) Update() Response {
	var r Response
	result := database.Database.Save(&a)
	if result.Error != nil {
		log.Println("an error occurred while updating the admin to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(a)
}

func (a Admin) Delete() Response {
	var r Response
	a.IsActive = false
	result := database.Database.Model(Admin{}).Where("person_id = ?", a.PersonId).Save(&a)
	if result.Error != nil {
		log.Println("an error occurred while deleting the admin to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(a)
}

func (a Admin) Get() Response {
	var r Response
	result := database.Database.Find(&a, "person_id = ? and is_active = ?", a.PersonId, true)
	if result.Error != nil {
		log.Println("an error occurred while getting the admin to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(a)
}

func (a Admin) List() Response {
	var r Response
	var admins []Admin
	result := database.Database.Find(&admins, "is_active = ?", true)
	if result.Error != nil {
		log.Println("an error occurred while listing the admins to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(admins)
}

func (a Admin) GetAdmin() Admin {
	result := database.Database.Find(&a, "person_id = ?", a.PersonId)
	if result.Error != nil {
		log.Println("ERROR:", result.Error.Error())
	}
	return a
}
