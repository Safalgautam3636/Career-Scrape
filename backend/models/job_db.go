package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// you gotta run this on your db first: CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
// https://stackoverflow.com/questions/36486511/how-do-you-do-uuid-in-golangs-gorm
type JobDB struct {
	gorm.Model
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	JobTitle           string    `json:"job_title" binding:"required" gorm:"uniqueIndex:idx_title_company_location"`
	JobLink            string    `json:"job_link" binding:"required"`
	JobLocation        string    `json:"job_location" binding:"required" gorm:"uniqueIndex:idx_title_company_location"`
	CompanyName     string    `json:"company_name" binding:"required" gorm:"uniqueIndex:idx_title_company_location"`
	JobType         string    `json:"job_type"`   // fulltime/contract
	CompanyDomain   string    `json:"company_domain"` //tech consulting manufacturing...
	JobLevel        string    `json:"job_level"`      //entry level/senior/architect...
	CompanyLink     string    `json:"company_link"`
	Description     string    `json:"description" binding:"required"`
	PulledTimeStamp time.Time `json:"pulled_date" gorm:"default:CURRENT_TIMESTAMP"`
	ExactDate       time.Time `json:"exact_date"` // actual date
	JobPosted       string    `json:"job_posted"` // 4 hrs ago...
	// CreatedAt       time.Time `json:"created_at,omitempty"`
	// UpdatedAt       time.Time `json:"updated_at,omitempty"`
	// DeletedAt       time.Time `json:"deleted_at,omitempty"`
}

