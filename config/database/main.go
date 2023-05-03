package database

import (
	"log"

	"github.com/daddydemir/VIKE-Backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitConnect() {
	dsn := config.Get("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	Database = db
}
