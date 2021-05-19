package models



type Voucher struct {
	Amount      float64   `json:"amount"`
	Hash        string    `json:"hash" gorm:"primary_key"`
	Status      string    `json:"status" gorm:"size:255;"`
	Code      string    `json:"code" gorm:"size:255;"`
	Batch      string    `json:"batch" gorm:"size:255;"`
	GeneratedBy string    `json:"generated_by" gorm:"size:255;"`
	CreatedAt int64 `json:"created_at" `
	UpdatedAt int64 `json:"updated_at" `
}

type VBatch struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Timeline        string    `json:"timeline" gorm:"size:255;"`
	Owner        string    `json:"owner" gorm:"size:255;"`
}
type VoucherPayload struct{
	Code        string    `json:"code"`
	MSISDN        string    `json:"msisdn"`
}

type VoucherCodeCheckPayload struct{
	Amount      float64   `json:"amount"`
	Status      string   `json:"status"`
}
