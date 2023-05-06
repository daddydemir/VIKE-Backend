package models

import (
	"github.com/daddydemir/VIKE-Backend/config/database"
	"log"
)

type Customer struct {
	Person
	IsVerified bool
}

func (c Customer) Add() Response {
	var r Response
	//TODO: use password hash
	result := database.Database.Create(&c)
	if result.Error != nil {
		log.Println("an error occurred while saving the customer to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(c)
}

func (c Customer) Update() Response {
	var r Response
	result := database.Database.Save(&c)
	if result.Error != nil {
		log.Println("an error occurred while updating the customer to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(c)
}

func (c Customer) Delete() Response {
	var r Response
	c.IsActive = false
	result := database.Database.Save(&c)
	if result.Error != nil {
		log.Println("an error occurred while deleting the customer to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(c)
}

func (c Customer) Get() Response {
	var r Response
	result := database.Database.Find(&c, "person_id = ? and is_active = ?", c.PersonId, true)
	if result.Error != nil {
		log.Println("an error occurred while getting the customer to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(c)
}

func (c Customer) List() Response {
	var r Response
	var customers []Customer
	result := database.Database.Find(&customers, "is_active = ? and is_verified = ?", true, true)
	if result.Error != nil {
		log.Println("an error occurred while listing the customers to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(customers)
}

// TODO: find method called and has returned success but data is empty! check this problem
