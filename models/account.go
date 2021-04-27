package models

import "time"

type Account struct {
	MSISDN    string    `json:"msisdn" gorm:"primary_key;unique;not null"`
	YOB       string    `json:"yob" gorm:"type:varchar(10);not null"`
	Balance   float64   `json:"balance"`
	RefCode   string    `json:"refcode" gorm:"size:255;"`
	Status    string    `json:"status" gorm:"size:255;"`
	CreatedAt time.Time `json:"created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" sql:"DEFAULT:current_timestamp"`
}