package models
import (
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
