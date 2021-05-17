package models


type Account struct {
	MSISDN    string    `json:"msisdn" gorm:"primary_key;unique;not null"`
	YOB       string    `json:"yob" gorm:"type:varchar(10);not null"`
	Balance   float64   `json:"balance"`
	RefCode   string    `json:"refcode" gorm:"size:255;"`
	Status    string    `json:"status" gorm:"size:255;"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime"`
}

