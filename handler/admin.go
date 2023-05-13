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

func adminAdd(w http.ResponseWriter, r *http.Request) {
	var admin mapper.AdminMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body Can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &admin)
	if err != nil {
		log.Println("Request Body can't parse to adminMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(service.AddExecute(admin.ToAdmin()))
	if err != nil {
		log.Println(err)
	}
}

func adminUpdate(w http.ResponseWriter, r *http.Request) {
	var admin mapper.AdminMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body Can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &admin)
	if err != nil {
		log.Println("Request Body can't parse to adminMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(service.UpdateExecute(admin.ToAdmin()))
	if err != nil {
		log.Println(err)
	}
}

func adminDelete(w http.ResponseWriter, r *http.Request) {
	var admin mapper.AdminMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body Can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &admin)
	if err != nil {
		log.Println("Request Body can't parse to adminMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(service.DeleteExecute(admin.ToAdmin()))
	if err != nil {
		log.Println(err)
	}
}

func adminGet(w http.ResponseWriter, r *http.Request) {
	var admin mapper.AdminMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body Can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &admin)
	if err != nil {
		log.Println("Request Body can't parse to adminMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(service.GetExecute(admin.ToAdmin()))
	if err != nil {
		log.Println(err)
	}
}

func adminList(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(service.ListExecute(models.Admin{}))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
}
