package cache

import (
	"context"
	"github.com/daddydemir/VIKE-Backend/config/cache"
	"github.com/daddydemir/VIKE-Backend/pkg/models"
	"github.com/google/uuid"
	"log"
	"time"
)

func ReadToken(sessionId, token string) bool {
	check := cache.Rdb.Ping(context.Background()).Err()
	if check != nil {
		log.Println("Cache server is unreachable :", check.Error())
		id, _ := uuid.Parse(sessionId)
		return models.Session{SessionId: id}.Get().Success
	}
	value := cache.Rdb.Get(context.Background(), sessionId)
	if value.Err() != nil {
		log.Println("key doesn't exists :", value.Err())
		return false
	}
	if value.Val() == token {
		return true
	} else {
		log.Println("Cache server is returned :", value.Val())
		return false
	}
}

func WriteToken(session models.Session) bool {
	check := cache.Rdb.Ping(context.Background()).Err()
	if check != nil {
		log.Println("Cache server is unreachable :", check.Error())
		return false
	}
	response := cache.Rdb.Set(context.Background(), session.SessionId.String(), session.Token, session.ExDate.Sub(time.Now()))
	if response.Err() != nil {
		log.Println("an error occurred has write to cache server :", response.Err())
		return false
	} else {
		return true
	}
}
