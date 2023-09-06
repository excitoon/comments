package v1

import "net/http"
import "strconv"

import "github.com/gin-gonic/gin"

import "auth"
import "crud"
import "models"

func GetUsers(c *gin.Context) {
	var currentUser = auth.Middleware.IdentityHandler(c).(*models.User)

	var users = crud.GetUsers()

	/// TODO pagination

	/// TODO https://stackoverflow.com/questions/17306358/removing-fields-from-struct-or-hiding-them-in-json-response
	if *currentUser.IsAdmin == false {
		for i := range users {
			users[i].Password = nil
			users[i].IsAdmin = nil
		}
	}

	c.IndentedJSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	var currentUser = auth.Middleware.IdentityHandler(c).(*models.User)

	var userIdStr = c.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		/// TODO common error handler
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	var user = crud.GetUser(uint(userId))
	if user == nil {
		/// TODO common error handler
		c.IndentedJSON(http.StatusNotFound, "")
		return
	}

	/// TODO https://stackoverflow.com/questions/17306358/removing-fields-from-struct-or-hiding-them-in-json-response
	if *currentUser.IsAdmin == false {
		user.Password = nil
		user.IsAdmin = nil
	}

	c.IndentedJSON(http.StatusOK, user)
}
