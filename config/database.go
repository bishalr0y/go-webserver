package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func ConnectToDb() *gorm.DB {
	dsn := "user=admin password=admin dbname=my_db port=5432 sslmode=disable TimeZone=Asia/Kolkata host=localhost"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
