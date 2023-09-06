package crud

import "db"
import "models"

func GetComments() []models.Comment {
	var comments []models.Comment

	db.DB.Find(&comments)

	return comments
}

func GetComment(commentId uint) *models.Comment {
	var comment models.Comment

	err := db.DB.Take(&comment, "Id = ?", commentId).Error
	if err != nil {
		return nil
	}

	return &comment
}
