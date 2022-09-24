package database

import (
	"os"
	"sync"

	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func CreateConnection() {
	var conf dbConfig

	if os.Getenv("APP_ENV") == "test" {
		conf = dbConfig{
			User: os.Getenv("DB_USER_TEST"),
			Pass: os.Getenv("DB_PASS_TEST"),
			Host: os.Getenv("DB_HOST_TEST"),
			Port: os.Getenv("DB_PORT_TEST"),
			Name: os.Getenv("DB_NAME_TEST"),
		}
	} else {
		conf = dbConfig{
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Name: os.Getenv("DB_NAME"),
		}
	}

	mysql := mysqlConfig{
		dbConfig: conf,
	}

	once.Do(func() {
		mysql.Connect()
	})
}

func GetConnection() *gorm.DB {
	if db == nil {
		CreateConnection()
	}

	return db
}
