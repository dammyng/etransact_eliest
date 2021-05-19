package handler

import (
	"eliest/helpers"
	"eliest/models"
	"encoding/json"
	"errors"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

func (handler *EliestHandler) Fund(w http.ResponseWriter, r *http.Request) {
	var reg models.FundingPayload


	err := json.NewDecoder(r.Body).Decode(&reg)
	defer r.Body.Close()
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	_, err = valid.ValidateStruct(reg)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error())
		return	
	}

	user, err := handler.Db.FindAccount(&models.Account{MSISDN: reg.MSISDN})
	if user == nil {
		helpers.RespondWithError(w, http.StatusNotFound, UserNotFound)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		helpers.RespondWithError(w, http.StatusNotFound, UserNotFound)
		return
	}
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, GeneralServiceError)
		return
	}

	err = handler.Db.UpdateUser(user, &models.Account{Balance: user.Balance + reg.Amount})
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, GeneralServiceError)
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, nil)

}



func (handler *EliestHandler) RechargeVoucher(w http.ResponseWriter, r *http.Request) {

	var winPayload models.VoucherCallback

	err := json.NewDecoder(r.Body).Decode(&winPayload)
	defer r.Body.Close()
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	a := []rune(winPayload.Code)
	pin := string(a[0:3])
	serial := string(a[3:7])
	validator := pin + serial
	hash := helpers.VoucherHash(validator, serial)

	win, err := handler.Db.FindVoucher(&hash)
	if errors.Is(err, gorm.ErrRecordNotFound)  {
		helpers.RespondWithError(w, http.StatusNotFound, "Invalid or used voucher")
		return
	}
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, GeneralServiceError)
			return
	}
	if win.Status != "used"  {
		user, err := handler.Db.FindAccount(&models.Account{MSISDN: winPayload.MSISDN})
		if err != nil {
			helpers.RespondWithError(w, http.StatusNotFound, UserNotFound)
			return
		}

		err = handler.Db.UpdateVoucherMap(win, map[string]interface{}{"status": "used"})
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, GeneralServiceError)
			return
		}
		
		err = handler.Db.UpdateUser(user, &models.Account{Balance: user.Balance + win.Amount})
		//create transaction for agent
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, GeneralServiceError)
			return
		}
		helpers.RespondWithJSON(w, http.StatusOK, "You have successfully transferred your winning")
		return
	} else {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid voucher code")
	}

}