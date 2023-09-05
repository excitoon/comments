package v1

import "net/http"

import "github.com/gin-gonic/gin"

import "crud"

func GetComments(c *gin.Context) {
	var comments = crud.GetComments()

	/// TODO pagination

	c.IndentedJSON(http.StatusOK, comments)
}
