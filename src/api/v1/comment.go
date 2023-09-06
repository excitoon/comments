package v1

import "net/http"
import "strconv"

import "github.com/gin-gonic/gin"

import "crud"

func GetComments(c *gin.Context) {
	var comments = crud.GetComments()

	/// TODO pagination

	c.IndentedJSON(http.StatusOK, comments)
}

func GetComment(c *gin.Context) {
	var commentIdStr = c.Param("commentId")
	commentId, err := strconv.ParseUint(commentIdStr, 10, 64)
	if err != nil {
		/// TODO common error handler
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	var comment = crud.GetComment(uint(commentId))
	if comment == nil {
		/// TODO common error handler
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	c.IndentedJSON(http.StatusOK, comment)
}
