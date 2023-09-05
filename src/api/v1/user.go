package v1

import "net/http"

import "github.com/gin-gonic/gin"

import "crud"


func GetUsers(c *gin.Context) {
    var users = crud.GetUsers()
    c.IndentedJSON(http.StatusOK, users)
}
