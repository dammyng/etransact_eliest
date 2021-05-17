package models



type Vouchers struct {
	Amount      float64   `json:"amount"`
	Hash        string    `json:"hash" gorm:"primary_key"`
	Status      string    `json:"status" gorm:"size:255;"`
	GeneratedBy string    `json:"generated_by" gorm:"size:255;"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime"`
}

type VoucherPayload struct{
	Code        string    `json:"code"`
	MSISDN        string    `json:"msisdn"`
}

type VoucherCodeCheckPayload struct{
	Amount      float64   `json:"amount"`
	Status      string   `json:"status"`
}
