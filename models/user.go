package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `json:"ID"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
