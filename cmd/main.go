package main

import (
	"log"

	"github.com/daddydemir/VIKE-Backend/config"
	"github.com/daddydemir/VIKE-Backend/config/database"
)

func main() {
	log.Println("Server started at http://localhost" + config.Get("PORT"))
	database.InitConnect()
	database.TableCreate()
}
