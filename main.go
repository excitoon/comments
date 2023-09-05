package main

import "net/http"
import "time"

import "github.com/gin-gonic/gin"
import "gorm.io/gorm"
import "gorm.io/driver/postgres"


type User struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
    Password string
}

type Comment struct {
    Id int `json:"id"`
    UserId string `json:"id_user"`
    Text string `json:"txt"`
}


func getUsers(c *gin.Context) {
    var users = []User{
        {Id: 1, Name: "The Administrator", Email: "admin@example.com"},
    }
    c.IndentedJSON(http.StatusOK, users)
}

func main() {
    dsn := "host=db user=postgres password=postgres dbname=database port=5432 sslmode=disable TimeZone=Europe/Moscow"
    conn, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    db, _ := conn.DB()
    db.SetMaxIdleConns(10)
    db.SetMaxOpenConns(100)
    db.SetConnMaxLifetime(time.Hour)

    router := gin.Default()
    router.GET("/api/v1/users", getUsers)

    router.Run("0.0.0.0:80")
}
