package main

import "github.com/gin-gonic/gin"

import "api"
import "auth"


func main() {
    router := gin.Default()

    group := router.Group("/")

    authenticatedGroup := router.Group("/")
    authenticatedGroup.Use(auth.Middleware.MiddlewareFunc())

    api.AddRoutes(group, authenticatedGroup)

    router.Run("0.0.0.0:80")
}
