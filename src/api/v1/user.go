package v1

import "net/http"

import "github.com/gin-gonic/gin"

import "auth"
import "crud"
import "models"


func GetUsers(c *gin.Context) {
    var currentUser = auth.Middleware.IdentityHandler(c).(*models.User)

    var users = crud.GetUsers()

    /// TODO https://stackoverflow.com/questions/17306358/removing-fields-from-struct-or-hiding-them-in-json-response
    if *currentUser.IsAdmin == false {
        for i, _ := range users {
            users[i].Password = nil;
            users[i].IsAdmin = nil;
        }
    }

    c.IndentedJSON(http.StatusOK, users)
}
