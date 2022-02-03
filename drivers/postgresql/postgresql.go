package postgresql

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	User string
}

func (c *DBConfig) InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%v", c.User)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// move to fatal so it can panic
		log.Println(err)
	}

	return db
}
