package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// you gotta run this on your db first: CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
// https://stackoverflow.com/questions/36486511/how-do-you-do-uuid-in-golangs-gorm
type Job struct {
	gorm.Model
	ID              uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Role            string    `json:"role"`
	Link            string    `json:"link"`
	Address         string    `json:"address"`
	CompanyName     string    `json:"company_name"`
	DatePosted      time.Time `json:"posted_date"`
	JobType         string    `json:"job_type"` //remote/onsite/hybrid
	Description     string    `json:"description"`
	Salary          string    `json:"salary"`
	PulledTimeStamp time.Time `json:"pulled_date" gorm:"default:CURRENT_TIMESTAMP"`
	ExperienceLevel string    `json:"experience_level"`
	// CreatedAt       time.Time `json:"created_at,omitempty"`
	// UpdatedAt       time.Time `json:"updated_at,omitempty"`
	// DeletedAt       time.Time `json:"deleted_at,omitempty"`
}