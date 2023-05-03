package main

import (
	"log"

	"github.com/daddydemir/VIKE-Backend/config"
)

func main() {
	log.Println("Server started at http://localhost" + config.Get("PORT"))
}
