package api

import "github.com/gin-gonic/gin"

import (
	"api.v1"
	"env"
)

func AddRoutes(group *gin.RouterGroup, authenticatedGroup *gin.RouterGroup) {
	api := group.Group(env.Endpoint)
	authenticatedApi := authenticatedGroup.Group(env.Endpoint)

	v1.AddRoutes(api, authenticatedApi)
}
