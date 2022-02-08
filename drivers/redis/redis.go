package redis

import (
	"github.com/go-redis/redis/v8"
)

type DBConfig struct {
	User string
}

func (c *DBConfig) InitDB() *redis.Client {
	redisOpts := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	db := redis.NewClient(redisOpts)

	return db
}
