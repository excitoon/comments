package main

import "github.com/gin-gonic/gin"

import "api.v1"


func main() {
    router := gin.Default()
    router.GET("/api/v1/users", v1.GetUsers)

    router.Run("0.0.0.0:80")
}
