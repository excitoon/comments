package db

import "fmt"
import "time"

import "gorm.io/gorm"
import "gorm.io/driver/postgres"

import "env"

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Moscow",
		env.DatabaseHost,
		env.DatabaseUser,
		env.DatabasePass,
		env.DatabaseName,
		env.DatabasePort,
	)
	DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	database, _ := DB.DB()
	database.SetMaxIdleConns(10)           /// TODO to config
	database.SetMaxOpenConns(100)          /// TODO to config
	database.SetConnMaxLifetime(time.Hour) /// TODO to config
}
