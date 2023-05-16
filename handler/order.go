package handler

import (
	"encoding/json"
	"github.com/daddydemir/VIKE-Backend/pkg/mapper"
	"github.com/daddydemir/VIKE-Backend/pkg/models"
	"github.com/daddydemir/VIKE-Backend/pkg/service"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
)

func orderAdd(w http.ResponseWriter, r *http.Request) {
	var order mapper.OrderMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &order)
	if err != nil {
		log.Println("Request Body can't parse to orderMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// todo: i need created user id {get token parsing}
	err = json.NewEncoder(w).Encode(service.AddExecute(order.ToOrder(uuid.New())))
	if err != nil {
		log.Println(err)
	}
}

func orderUpdate(w http.ResponseWriter, r *http.Request) {
	var order mapper.OrderMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &order)
	if err != nil {
		log.Println("Request Body can't parse to orderMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// todo: i need created user id {get token parsing}
	err = json.NewEncoder(w).Encode(service.UpdateExecute(order.ToOrder(uuid.New())))
	if err != nil {
		log.Println(err)
	}
}

func orderDelete(w http.ResponseWriter, r *http.Request) {
	var id mapper.IdMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &id)
	if err != nil {
		log.Println("Request Body can't parse to orderMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// todo: i need created user id {get token parsing}
	err = json.NewEncoder(w).Encode(service.DeleteExecute(orderHelper(id)))
	if err != nil {
		log.Println(err)
	}
}

func orderGet(w http.ResponseWriter, r *http.Request) {
	var id mapper.IdMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &id)
	if err != nil {
		log.Println("Request Body can't parse to orderMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// todo: i need created user id {get token parsing}
	err = json.NewEncoder(w).Encode(service.GetExecute(id.ToOrder()))
	if err != nil {
		log.Println(err)
	}
}

func orderList(w http.ResponseWriter, r *http.Request) {
	var id mapper.IdMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &id)
	if err != nil {
		log.Println("Request Body can't parse to orderMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(service.ListExecute(id.ToOrder()))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func orderHelper(idMapper mapper.IdMapper) models.Order {
	var order models.Order
	order.OrderId = idMapper.Parse()
	return order.GetOrder()
}
