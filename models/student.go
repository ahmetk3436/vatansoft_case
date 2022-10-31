package models

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}
