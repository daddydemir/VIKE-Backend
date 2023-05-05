package models

import "fmt"

type Customer struct {
	Person
	IsVerified bool
}

func (c Customer) Add() Response {
	var r Response
	r.Data = c
	// save to database!
	fmt.Println("added.")
	return r
}

func (c Customer) Update() Response {
	return Response{}
}

func (c Customer) Delete() Response {
	return Response{}
}

func (c Customer) Get() Response {
	var r Response
	return r
}

func (c Customer) List() Response {
	var r Response
	return r
}
