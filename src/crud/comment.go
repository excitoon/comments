package crud

import "db"
import "models"

func GetComments() []models.Comment {
	var comments []models.Comment

	db.DB.Find(&comments)

	return comments
}
