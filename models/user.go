package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int    `json:"age"`
	City      string `json:"city"`
}
