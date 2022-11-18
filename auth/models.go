package auth

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string
	FirstName string
	LastName  string
	Password  string
}
