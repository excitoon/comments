package models

type Comment struct {
	Id     int    `json:"id" gorm:"id"`
	UserId string `json:"id_user" gorm:"id_user"`
	Text   string `json:"txt" gorm:"txt"`
}
