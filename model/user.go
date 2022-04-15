package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int    `json:"_id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
