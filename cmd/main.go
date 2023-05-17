package main

import (
	"github.com/daddydemir/VIKE-Backend/config/cache"
	"github.com/daddydemir/VIKE-Backend/handler"
	"log"
	"net/http"

	"github.com/daddydemir/VIKE-Backend/config"
	"github.com/daddydemir/VIKE-Backend/config/database"
)

func main() {
	log.Println("Server started at http://localhost" + config.Get("PORT"))
	database.InitConnect()
	cache.RedisClient()
	server := &http.Server{
		Addr:    config.Get("PORT"),
		Handler: handler.Route(),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server is down :", err)
	}
}
