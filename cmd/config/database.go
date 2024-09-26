package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=12345 dbname=gogrpc port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Gagal Connect")
	} 
	return database
}