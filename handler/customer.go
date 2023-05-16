package handler

import (
	"encoding/json"
	"github.com/daddydemir/VIKE-Backend/pkg/mapper"
	"github.com/daddydemir/VIKE-Backend/pkg/models"
	"github.com/daddydemir/VIKE-Backend/pkg/service"
	"io/ioutil"
	"log"
	"net/http"
)

func customerAdd(w http.ResponseWriter, r *http.Request) {
	var model mapper.CustomerMapper
	var verify mapper.VerifyMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body can't parse to json :", err)
	}
	err = json.Unmarshal(body, &model)
	if err != nil {
		log.Println("Request Body can't parse to customerMapper :", err)
	}
	w.WriteHeader(200)

	verify.Code = "133713"
	resp := service.AddExecute(model.ToCustomer())
	service.AddExecute(verify.ToVerify(model.ToCustomer().PersonId))
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
	}
}

func customerUpdate(w http.ResponseWriter, r *http.Request) {
	var customer mapper.CustomerMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &customer)
	if err != nil {
		log.Println("Request Body can't parse to customerMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(service.UpdateExecute(customer.ToCustomer()))
	if err != nil {
		log.Println(err)
	}
}

func customerDelete(w http.ResponseWriter, r *http.Request) {
	var id mapper.IdMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body Can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &id)
	if err != nil {
		log.Println("Request Body can't parse to customerMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(service.DeleteExecute(customerHelper(id)))
	if err != nil {
		log.Println(err)
	}
}

func customerGet(w http.ResponseWriter, r *http.Request) {
	var id mapper.IdMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &id)
	if err != nil {
		log.Println("Request Body can't parse to customerMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(service.GetExecute(id.ToCustomer()))
	if err != nil {
		log.Println(err)
	}
}

func customerList(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(service.ListExecute(models.Customer{}))
	if err != nil {
		log.Println(err)
	}
}

func customerHelper(idMapper mapper.IdMapper) models.Customer {
	var customer models.Customer
	customer.PersonId = idMapper.Parse()
	return customer.GetCustomer()
}
