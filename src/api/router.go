package api

import "github.com/gin-gonic/gin"

import "api.v1"

func AddRoutes(group *gin.RouterGroup, authenticatedGroup *gin.RouterGroup) {
	api := group.Group("/api")
	authenticatedApi := authenticatedGroup.Group("/api")

	v1.AddRoutes(api, authenticatedApi)
}
