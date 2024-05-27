package models

import "time"

type Job struct{
	ID uint `json:"id" gorm:"primary_key"`
	Role string `json:"role"`
	Link string `json:"link"`
	Address string `json:"address"`
	CompanyName string `json:"company_name"`
	DatePosted time.Time `json:"start_date"`
	AboutJob string `json:"about_job"`
	Qualifications string `json:"qualifications"`
	Responsibility string `json:"responsibility"`
	Benefits string `json:"benefits"`
}