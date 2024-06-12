package models

import "gorm.io/gorm"


type User struct{
	gorm.Model
	ID string `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"size:255;not null;unique"`
	Email string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"unique;not null"`
	IsAdmin  bool   `json:"isAdmin" gorm:"default:false"`
}