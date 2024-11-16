package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	// ID string `json:"id" gorm:"primary_key"`
	Username       string         `json:"username" binding:"required" gorm:"size:255;not null;unique"`
	Email          string         `json:"email" binding:"required" gorm:"unique;not null"`
	Password       string         `json:"password" binding:"required" gorm:"not null"`
	Followers      pq.StringArray `json:"followers" gorm:"type:text[];default:'{}'"`
	Following      pq.StringArray `json:"following" gorm:"type:text[];default:'{}'"`
	InterestedJobs pq.StringArray `json:"interestedJobs" gorm:"type:text[];default:'{}'"`
	IsAdmin        bool           `json:"isAdmin" gorm:"default:false"`
}
type LoginUser struct {
	Username string `json:"username" binding:"required" gorm:"size:255;not null;unique"`
	Password string `json:"password" binding:"required" gorm:"unique;not null"`
}

type ReturnUser struct {
	Username string `json:"username" gorm:"size:255;not null;unique"`
	Email    string `json:"email" gorm:"unique;not null"`
	IsAdmin  bool   `json:"isAdmin" gorm:"default:false"`
}
type VerifyRequestBody struct {
	Token string `json:"token"`
}
