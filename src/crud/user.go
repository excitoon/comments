package crud

import "db"
import "models"


func GetUsers() []models.User{
    var users []models.User

    db.DB.Find(&users)

    return users
}

func GetUserByName(name string) *models.User {
    var user models.User

    db.DB.Take(&user, "Name = ?", name)

    return &user
}
