package models

import (
	"github.com/daddydemir/VIKE-Backend/config/database"
	"log"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderId      uuid.UUID
	AdvertId     uuid.UUID
	CustomerId   uuid.UUID
	NumberOfUser int
	Status       string
	Price        float32
	ExpireDate   time.Time
}

func (o Order) Add() Response {
	var r Response
	result := database.Database.Create(&o)
	if result.Error != nil {
		log.Println("an error occurred while saving the order to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(o)
}

func (o Order) Update() Response {
	var r Response
	result := database.Database.Save(&o)
	if result.Error != nil {
		log.Println("an error occurred while updating the order to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(o)
}

func (o Order) Delete() Response {
	var r Response
	o.Status = "CL"
	result := database.Database.Model(Order{}).Where("order_id = ?", o.OrderId).Save(&o)
	if result.Error != nil {
		log.Println("an error occurred while deleting the order to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(o)
}

func (o Order) Get() Response {
	var r Response
	result := database.Database.Find(&o, "order_id = ?", o.OrderId)
	if result.Error != nil {
		log.Println("an error occurred while getting the order to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(o)
}

func (o Order) List() Response {
	var r Response
	var orders []Order
	result := database.Database.Find(&orders, "customer_id = ?", o.CustomerId)
	if result.Error != nil {
		log.Println("an error occurred while listing the orders to db :", result.Error)
		return r.ErrorResponse(result.Error.Error())
	}
	return r.SuccessResponse(orders)
}

func (o Order) GetOrder() Order {
	result := database.Database.Find(&o, "order_id = ?", o.OrderId)
	if result.Error != nil {
		log.Println("ERROR:", result.Error.Error())
	}
	return o
}
