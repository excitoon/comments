package models

type User struct {
	Id       int     `json:"id" gorm:"column:id"`
	Name     string  `json:"name" gorm:"column:name"`
	Email    string  `json:"email" gorm:"column:email"`
	Password *string `json:"password,omitempty" gorm:"column:password"`
	IsAdmin  *bool   `json:"is_admin,omitempty" gorm:"column:is_admin"`
}

func (u *User) TableName() string {
	return "schema.users"
}
