package redis

import (
	"github.com/go-redis/redis/v8"
)

type DBConfig struct {
	User string
}

func (c *DBConfig) InitDB() *redis.Client {
	// dsn := fmt.Sprintf("localhost:6379")

	redisOpts := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	db := redis.NewClient(redisOpts)
	// if err != nil {
	// 	// move to fatal so it can panic
	// 	log.Println(err)
	// }

	return db
}
