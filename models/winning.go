package models

import "time"


type Winnings struct {
	Amount      float64   `json:"amount"`
	Hash        string    `json:"hash" gorm:"primary_key"`
	Status      string    `json:"status" gorm:"size:255;"`
	GeneratedBy string    `json:"generated_by" gorm:"size:255;"`
	CreatedAt   time.Time `json:"created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" sql:"DEFAULT:current_timestamp"`
}

type WinPayload struct{
	Code        string    `json:"code"`
	MSISDN        string    `json:"msisdn"`
}

type WinCodeCheckPayload struct{
	Amount      float64   `json:"amount"`
	Status      string   `json:"status"`
}

type TransferredCallback struct{
	Code        string    `json:"code"`
	MSISDN        string    `json:"msisdn"`
	Agent        string    `json:"agent"`
}
