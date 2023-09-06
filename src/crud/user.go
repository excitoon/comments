package crud

import "db"
import "models"

func GetUsers() []models.User {
	var users []models.User

	db.DB.Find(&users)

	return users
}

func GetUser(userId uint) *models.User {
	var user models.User

	err := db.DB.Take(&user, "Id = ?", userId).Error
	if err != nil {
		return nil
	}

	return &user
}

func GetUserByName(name string) *models.User {
	var user models.User

	err := db.DB.Take(&user, "Name = ?", name).Error
	if err != nil {
		return nil
	}

	return &user
}
