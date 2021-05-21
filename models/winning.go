package models



type Winnings struct {
	Amount      float64   `json:"amount"`
	Hash        string    `json:"hash" gorm:"primary_key"`
	Status      string    `json:"status" gorm:"size:255;"`
	Code      string    `json:"code" gorm:"size:255;"`
	GeneratedBy string    `json:"generated_by" gorm:"size:255;"`
	CreatedAt int64 `json:"created_at" `
	UpdatedAt int64 `json:"updated_at"`
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


type VoucherCallback struct{
	Code        string    `json:"code"`
	MSISDN        string    `json:"msisdn"`
}