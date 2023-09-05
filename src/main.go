package main

import "fmt"

import "github.com/gin-gonic/gin"

import "api"
import "auth"
import "env"

func main() {
	router := gin.Default()

	group := router.Group("/")

	authenticatedGroup := router.Group("/")
	authenticatedGroup.Use(auth.Middleware.MiddlewareFunc())

	api.AddRoutes(group, authenticatedGroup)

	router.Run(fmt.Sprintf("%s:%d", env.Host, env.Port))
}
