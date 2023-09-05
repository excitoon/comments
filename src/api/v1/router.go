package v1

import "github.com/gin-gonic/gin"

import "auth"

func AddRoutes(group *gin.RouterGroup, authenticatedGroup *gin.RouterGroup) {
	v1 := group.Group("/v1")
	authenticatedV1 := authenticatedGroup.Group("/v1")

	v1.POST("/auth", auth.Middleware.LoginHandler)

	authenticatedV1.GET("/users", GetUsers)
	authenticatedV1.GET("/comments", GetComments)
}
