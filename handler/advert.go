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

func advertAdd(w http.ResponseWriter, r *http.Request) {
	var advert mapper.AdvertMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &advert)
	if err != nil {
		log.Println("Request Body can't parse to advertMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//todo: i need this request creator user's uuid
	err = json.NewEncoder(w).Encode(service.AddExecute(advert.ToAdvert(uuid.New())))
	if err != nil {
		log.Println(err)
	}
}

func advertUpdate(w http.ResponseWriter, r *http.Request) {
	var advert mapper.AdvertMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &advert)
	if err != nil {
		log.Println("Request Body can't parse to advertMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//todo: i need this request creator user's uuid
	err = json.NewEncoder(w).Encode(service.UpdateExecute(advert.ToAdvert(uuid.New())))
	if err != nil {
		log.Println(err)
	}
}

func advertDelete(w http.ResponseWriter, r *http.Request) {
	var id mapper.IdMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &id)
	if err != nil {
		log.Println("Request Body can't parse to advertMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(service.DeleteExecute(advertHelper(id)))
	if err != nil {
		log.Println(err)
	}
}

func advertGet(w http.ResponseWriter, r *http.Request) {
	var id mapper.IdMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &id)
	if err != nil {
		log.Println("Request Body can't parse to advertMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(service.GetExecute(id.ToAdvert()))
	if err != nil {
		log.Println(err)
	}
}

func advertList(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(service.ListExecute(models.Advert{}))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func advertHelper(idMapper mapper.IdMapper) models.Advert {
	var advert models.Advert
	advert.AdvertId = idMapper.Parse()
	return advert.GetAdvert()
}
