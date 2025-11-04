package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(redisUrl string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisUrl, // host:port
		// Password: "",        // set if you have a password
		DB: 0, // use default DB
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	log.Println("Connected to Redis. ")

	return rdb, nil
}
