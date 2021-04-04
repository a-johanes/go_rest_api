package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {

	dsn := "host=localhost user=postgres password=password dbname=go_rest_api port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	if err = database.AutoMigrate(&Book{}); err != nil {
		return err
	}

	DB = database
	return nil
}
