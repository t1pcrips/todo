package user

import "gorm.io/gorm"

type User struct {
	Email    string
	Password string
	gorm.Model
}
