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
	return

}
