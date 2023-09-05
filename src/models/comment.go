package models

type Comment struct {
	Id     int    `json:"id" gorm:"column:id"`
	UserId int    `json:"id_user" gorm:"column:id_user"`
	Text   string `json:"txt" gorm:"column:txt"`
}

func (c *Comment) TableName() string {
	return "schema.comments"
}
