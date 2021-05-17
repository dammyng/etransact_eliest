package models


type Agent struct {
	Id               string     `json:"id"  gorm:"primary_key"`
	Firstname        string    `json:"firstname" gorm:"size:255;"`
	Lastname         string    `json:"lastname" gorm:"size:255;"`
	Password         []byte    `json:"-"`
	Email            string    `json:"email" gorm:"size:255;unique"`
	Status           string    `json:"status" gorm:"size:255;"`
	Phone            string    `json:"phone" gorm:"size:255;unique"`
	State            string    `json:"state" gorm:"size:255;"`
	City            string    `json:"city" gorm:"size:255;"`
	Lg            string    `json:"lg" gorm:"size:255;"`
	Address            string    `json:"address" gorm:"size:255;"`
	RefCode          string    `json:"refcode" gorm:"size:255;"`
	Referrer         string    `json:"referrer" gorm:"size:255;"`
	AccountActivated bool      `json:"account_verified"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime"`
}


type Wallet struct {
	Id        string    `json:"id"  gorm:"primary_key"`
	Balance   float64   `json:"balance"`
	Owner     string    `json:"owner" gorm:"primary_key"`
	Status    string    `json:"status" gorm:"size:255;"`
	Class     string    `json:"class" gorm:"size:255;"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime"`
}

type Transaction struct {
	Id          string    `json:"id"  gorm:"primary_key"`
	Amount      float64   `json:"amount"`
	TRef        string    `json:"t_ref" gorm:"primary_key"`
	Account     string    `json:"account" gorm:"size:255;"`
	Status      string    `json:"status" gorm:"size:255;"`
	Description string    `json:"description" gorm:"size:255;"`
	Class       string    `json:"class" gorm:"size:255;"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime"`
}