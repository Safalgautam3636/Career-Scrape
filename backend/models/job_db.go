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
	JobTitle           string    `json:"job_title" gorm:"uniqueIndex:idx_title_company_location"`
	JobLink            string    `json:"job_link"`
	JobLocation        string    `json:"job_location" gorm:"uniqueIndex:idx_title_company_location"`
	JobPosted      string    `json:"job_posted"` // 4 hrs ago...
	CompanyName     string    `json:"company_name" gorm:"uniqueIndex:idx_title_company_location"`
	ExactDate       time.Time `json:"exact_date"` // actual date
	JobType         string    `json:"job_type"`   // fulltime/contract
	CompanyDomain   string    `json:"company_domain"`
	JobLevel        string    `json:"job_level"`
	CompanyLink     string    `json:"company_link"`
	Description     string    `json:"description"`
	PulledTimeStamp time.Time `json:"pulled_date" gorm:"default:CURRENT_TIMESTAMP"`
	// CreatedAt       time.Time `json:"created_at,omitempty"`
	// UpdatedAt       time.Time `json:"updated_at,omitempty"`
	// DeletedAt       time.Time `json:"deleted_at,omitempty"`
}

