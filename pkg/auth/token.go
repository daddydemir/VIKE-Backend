package auth

import (
	"fmt"
	"github.com/daddydemir/VIKE-Backend/config"
	"github.com/daddydemir/VIKE-Backend/pkg/cache"
	"github.com/daddydemir/VIKE-Backend/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

func GenerateToken(model models.Session) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["user"] = model.Email
	claims["session"] = model.SessionId
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7)

	tokenString, _ := token.SignedString([]byte(config.Get("JWT_KEY")))
	model.Token = tokenString
	model.Add()
	cache.WriteToken(model)
	return tokenString
}

func IsValid(data string) models.Response {
	var r models.Response
	token, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get("JWT_KEY")), nil
	})
	if err != nil {
		log.Println("an error occurred has token parsing :", err)
		return r.ErrorResponse(err.Error())
	}
	if !token.Valid {
		return r.ErrorResponse("unauthorized")
		// todo: detail error message will be add
	} else {
		return r.SuccessResponse(nil)
	}
}

func TokenParser(data string) string {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(data, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get("JWT_KEY")), nil
	})
	if err != nil {
		log.Println("an error occurred has token parsing :", err)
		return ""
	}
	sessionId := fmt.Sprintf("%v", claims["session"])
	return sessionId
}
