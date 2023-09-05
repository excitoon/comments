package models


type User struct {
    Id int `json:"id" gorm:"id"`
    Name string `json:"name" gorm:"name"`
    Email string `json:"email" gorm:"email"`
    Password *string `json:"password,omitempty" gorm:"password"`
    IsAdmin *bool `json:"is_admin,omitempty" gorm:"is_admin"`
}

func (u *User) TableName() string {
    return "schema.users"
}
