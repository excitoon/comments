package v1

import (
	"net/http"
	"strconv"
)

import "github.com/gin-gonic/gin"

import (
	"auth"
	"crud"
	"models"
)

func GetComments(c *gin.Context) {
	comments := crud.GetComments()

	/// TODO pagination

	c.IndentedJSON(http.StatusOK, comments)
}

func GetComment(c *gin.Context) {
	commentIdStr := c.Param("commentId")
	commentId, err := strconv.ParseUint(commentIdStr, 10, 64)
	if err != nil {
		/// TODO common error handler
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	comment := crud.GetComment(uint(commentId))
	if comment == nil {
		/// TODO common error handler
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	c.IndentedJSON(http.StatusOK, comment)
}

func GetUserComments(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		/// TODO common error handler
		/// TODO common int param code
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	comments := crud.GetUserComments(uint(userId))

	/// TODO pagination

	c.IndentedJSON(http.StatusOK, comments)
}

func GetUserComment(c *gin.Context) {
	commentIdStr := c.Param("commentId")
	commentId, err := strconv.ParseUint(commentIdStr, 10, 64)
	if err != nil {
		/// TODO common error handler
		/// TODO common int param code
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	userIdStr := c.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		/// TODO common error handler
		/// TODO common int param code
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	comment := crud.GetUserComment(uint(userId), uint(commentId))
	if comment == nil {
		/// TODO common error handler
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	c.IndentedJSON(http.StatusOK, comment)
}

func AddUserComment(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		/// TODO common error handler
		/// TODO common int param code
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	user := crud.GetUser(uint(userId))
	if user == nil {
		/// TODO common error handler
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	currentUser := auth.Middleware.IdentityHandler(c).(*models.User)
	if !*currentUser.IsAdmin {
		if user.Name != currentUser.Name { /// TODO check by Id
			/// TODO common error handler
			c.IndentedJSON(http.StatusForbidden, "")
			return
		}
	}

	var comment models.Comment
	c.BindJSON(&comment)

	comment.UserId = user.Id

	if ok := crud.AddUserComment(&comment); !ok {
		/// TODO common error handler
		c.IndentedJSON(http.StatusInternalServerError, "")
		return
	}

	c.IndentedJSON(http.StatusOK, comment)
}

func UpdateComment(c *gin.Context) {
	commentIdStr := c.Param("commentId")
	commentId, err := strconv.ParseUint(commentIdStr, 10, 64)
	if err != nil {
		/// TODO common error handler
		/// TODO common int param code
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	comment := crud.GetComment(uint(commentId))
	if comment == nil {
		/// TODO common error handler
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	user := crud.GetUser(comment.UserId)
	if user == nil {
		/// TODO common error handler
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	currentUser := auth.Middleware.IdentityHandler(c).(*models.User)
	if !*currentUser.IsAdmin {
		if user.Name != currentUser.Name { /// TODO check by Id
			/// TODO common error handler
			c.IndentedJSON(http.StatusForbidden, "")
			return
		}
	}

	var updateComment models.Comment
	c.BindJSON(&updateComment)

	if ok := crud.UpdateCommentText(comment.Id, updateComment.Text); !ok {
		/// TODO common error handler
		c.IndentedJSON(http.StatusInternalServerError, "")
		return
	}

	comment.Text = updateComment.Text
	c.IndentedJSON(http.StatusOK, comment)
}

func DeleteComment(c *gin.Context) {
	commentIdStr := c.Param("commentId")
	commentId, err := strconv.ParseUint(commentIdStr, 10, 64)
	if err != nil {
		/// TODO common error handler
		/// TODO common int param code
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	comment := crud.GetComment(uint(commentId))
	if comment == nil {
		/// TODO common error handler
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	user := crud.GetUser(comment.UserId)
	if user == nil {
		/// TODO common error handler
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	currentUser := auth.Middleware.IdentityHandler(c).(*models.User)
	if !*currentUser.IsAdmin {
		if user.Name != currentUser.Name { /// TODO check by Id
			/// TODO common error handler
			c.IndentedJSON(http.StatusForbidden, "")
			return
		}
	}

	if ok := crud.DeleteComment(comment.Id); !ok {
		/// TODO common error handler
		c.IndentedJSON(http.StatusInternalServerError, "")
		return
	}

	c.IndentedJSON(http.StatusOK, comment)
}
