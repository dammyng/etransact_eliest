package handler

import (
	"eliest/helpers"
	"eliest/models"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	valid "github.com/asaskevich/govalidator"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (handler *EliestHandler) Register(w http.ResponseWriter, r *http.Request) {
	
	var reg models.RegistrationPayload

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
	if user != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, DoubleRegistration)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		id, err := handler.Db.CreateAccount(&models.Account{
			MSISDN:    reg.MSISDN,
			YOB:       reg.YOB,
			Balance:   0.00,
			RefCode:   helpers.RandInt(6),
			Status:    "active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, GeneralServiceError)
			return
		}
		helpers.RespondWithJSON(w, http.StatusOK, id)
		return
	}
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, GeneralServiceError)
		return
	}
}

func (handler *EliestHandler) Details(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	msisdn := params["msisdn"]

	user, err := handler.Db.FindAccount(&models.Account{MSISDN: msisdn})
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
	helpers.RespondWithJSON(w, http.StatusOK, user)
	return
}
