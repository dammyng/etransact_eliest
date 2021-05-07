package models

import (
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
  }
type RegistrationPayload struct {
	MSISDN string `json:"msisdn" valid:"type(string)"`
	YOB    string `json:"yob" valid:"type(string)"`
}

type FundingPayload struct {
	Amount float64 `json:"amount" valid:"type(float64)"`
	MSISDN string  `json:"msisdn" valid:"type(string)"`
}

type GamePlayPayload struct {
	GameID string `json:"game_id" valid:"type(string)"`
	Guess string  `json:"guess" valid:"type(string)"`
	MSISDN string  `json:"msisdn" valid:"type(string)"`
}



type CoralDetailPayload struct {
	CustomerRef string `json:"customerRef"`
	MerchantId  string `json:"merchantId"`
}

type CoralDetailResponse struct {
	TraceId        string `json:"traceId"`
	CustomerName   string `json:"customerName"`
	DisplayMessage string `json:"displayMessage"`
	ResponseCode   string `json:"responseCode"`
}

type CoralNotifPayload struct {
	PassBackReference string `json:"passBackReference"`
	TraceId           string `json:"traceId"`
	CustomerRef       string `json:"customerRef"`
	PaymentReference  string `json:"paymentReference"`
	ResponseCode      string `json:"responseCode"`
	MerchantId        string `json:"merchantId"`
	MobileNumber      string `json:"mobileNumber"`
	Amount            float64  `json:"amount"`
	TransactionDate   string `json:"transactionDate"`
	ShortCode         string `json:"shortCode"`
	Currency          string `json:"currency"`
	Channel           string `json:"Channel"`
	Hash              string `json:"hash"`
}

type CoralNotifResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

func (cnp *CoralNotifPayload) HashValue() string {
	d := cnp.PaymentReference + cnp.CustomerRef + cnp.ResponseCode + cnp.MerchantId  + fmt.Sprintf("%.2f", cnp.Amount) + "|48fa5d5e-a5e8-4f52-8077-e52df2072443" 
	log.Println(d)

	return computeHmac256(d)
}

func computeHmac256(message string) string {
	ha := sha256.Sum256([]byte(message))
	return fmt.Sprintf("%x", ha)
}

