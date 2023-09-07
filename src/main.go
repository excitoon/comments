package main

import "fmt"

import "github.com/gin-gonic/gin"

import (
	"api"
	"auth"
	"env"
)

func main() {
	engine := gin.New()

	group := engine.Group("/")

	authenticatedGroup := engine.Group("/")
	authenticatedGroup.Use(auth.Middleware.MiddlewareFunc())

	api.AddRoutes(group, authenticatedGroup)

	engine.Run(fmt.Sprintf("%s:%d", env.Host, env.Port))
}
