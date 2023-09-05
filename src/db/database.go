package db

import "time"

import "gorm.io/gorm"
import "gorm.io/driver/postgres"


var DB *gorm.DB


func init() {
    dsn := "host=db user=postgres password=postgres dbname=database port=5432 sslmode=disable TimeZone=Europe/Moscow"
    DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    database, _ := DB.DB()
    database.SetMaxIdleConns(10)
    database.SetMaxOpenConns(100)
    database.SetConnMaxLifetime(time.Hour)
}
