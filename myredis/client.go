package myredis

import (
	"github.com/go-redis/redis/v7"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(dsn , pass string) RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: dsn, Password: pass, 
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return RedisClient{
		Client: client,
	}
}
