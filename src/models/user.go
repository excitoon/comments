package models


type User struct {
    Id int `json:"id" gorm:"id"`
    Name string `json:"name" gorm:"name"`
    Email string `json:"email" gorm:"email"`
    Password string `json:"-" gorm:"password"`
}

func (u *User) TableName() string {
    return "schema.users"
}
