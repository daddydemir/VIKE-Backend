package handler

import (
	"encoding/json"
	auth2 "github.com/daddydemir/VIKE-Backend/pkg/auth"
	"github.com/daddydemir/VIKE-Backend/pkg/encryption"
	"github.com/daddydemir/VIKE-Backend/pkg/mapper"
	"github.com/daddydemir/VIKE-Backend/pkg/models"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	var login mapper.LoginMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &login)
	if err != nil {
		log.Println("Request Body can't parse to loginMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	auth := check(login.Password, login.ToCustomer().Password)
	if auth {
		err = json.NewEncoder(w).Encode(auth2.GenerateToken(createSession(login.ToCustomer().Person, r)))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "message": "Mail or Password is wrong."})
	}
	if err != nil {
		log.Println(err)
	}
}

func adminLogin(w http.ResponseWriter, r *http.Request) {
	var login mapper.LoginMapper
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Request Body can't parse to json :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &login)
	if err != nil {
		log.Println("Request Body can't parse to loginMapper :", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	auth := check(login.Password, login.ToAdmin().Password)
	if auth {
		err = json.NewEncoder(w).Encode(auth2.GenerateToken(createSession(login.ToAdmin().Person, r)))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(map[string]interface{}{"success": false, "message": "Mail or Password is wrong."})
	}
	if err != nil {
		log.Println(err)
	}
}

func check(password string, hash string) bool {
	return encryption.CheckPass(hash, password)
}

func createSession(person models.Person, r *http.Request) models.Session {
	var session models.Session
	session.SessionId = uuid.New()
	session.Email = person.Email
	session.Ip = r.RemoteAddr
	session.CrDate = time.Now()
	session.ExDate = time.Now().Add(time.Hour * 24 * 7)
	return session
}
