package models

import "fmt"

import "env"

type User struct {
	Id       uint    `json:"id" gorm:"column:id;primaryKey"`
	Name     string  `json:"name" gorm:"column:name"`
	Email    string  `json:"email" gorm:"column:email"`
	Password *string `json:"password,omitempty" gorm:"column:password"`
	IsAdmin  *bool   `json:"is_admin,omitempty" gorm:"column:is_admin"`
}

func (u *User) TableName() string {
	return fmt.Sprintf("%s.users", env.DatabaseSchema)
}
