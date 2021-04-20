package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName" gorm:"size:100;not null"`
	LastName  string `json:"lastName" gorm:"size:100;not null"`
	Email     string `json:"email" gorm:"size:250;not null;unique"`
	Username  string `json:"username" gorm:"size:100;not null;unique;index"`
	Password  string `json:"password,omitempty" gorm:"size:300;not null"`
}
