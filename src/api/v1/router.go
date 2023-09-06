package v1

import "github.com/gin-gonic/gin"

import "auth"

func AddRoutes(group *gin.RouterGroup, authenticatedGroup *gin.RouterGroup) {
	v1 := group.Group("/v1")
	authenticatedV1 := authenticatedGroup.Group("/v1")

	v1.POST("/auth", auth.Middleware.LoginHandler)

	authenticatedV1.GET("/users", GetUsers)
	authenticatedV1.GET("/user/:userId", GetUser)

	authenticatedV1.GET("/user/:userId/comments", GetUserComments)
	authenticatedV1.GET("/user/:userId/comment/:commentId", GetUserComment)
	authenticatedV1.POST("/user/:userId/comment", AddUserComment)

	authenticatedV1.GET("/comments", GetComments)
	authenticatedV1.GET("/comment/:commentId", GetComment)
	authenticatedV1.PUT("/comment/:commentId", UpdateComment)
	authenticatedV1.DELETE("/comment/:commentId", DeleteComment)
}
