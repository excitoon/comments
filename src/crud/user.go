package crud

import "db"
import "models"


func GetUsers() []models.User{
    var users []models.User

    db.DB.Find(&users)

    return users
}
