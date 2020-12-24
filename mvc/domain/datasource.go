package domain

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbClient *gorm.DB
)

func init() {
	var err error
	dsn := "host=localhost user=primz password=primz@2207 dbname=primz port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	dbClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
