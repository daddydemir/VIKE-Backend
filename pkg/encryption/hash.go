package encryption

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPass(data string) string {
	pass, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		log.Println("an error occurred while encryption password :", err)
		return ""
	}
	return string(pass)
}

func CheckPass(encrypted, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(password))
	if err != nil {
		log.Println("an error occurred while password and hash comparing :", err)
		return false
	}
	return true
}
