package model

type User struct {
	ID       int64  `form:"id" gorm:"primary_key;not_null;auto_increment" json:"id"`
	Name     string `form:"name" gorm:"not_null" json:"name"`
	Password string `form:"password" gorm:"not_null" json:"password"`
	Email    string `form:"email" gorm:"not_null" json:"email"`
}
