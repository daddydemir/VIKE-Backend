package cache

import (
	"log"
	"strconv"

	"github.com/daddydemir/VIKE-Backend/config"
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func RedisClient() {
	db, err := strconv.Atoi(config.Get("REDIS_DB"))
	if err != nil {
		log.Fatal("Redis error:", err)
	}
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Get("REDIS_ADDR"),
		Password: config.Get("REDIS_PASS"),
		DB:       db,
	})
}
