package handler

import (
	"eliest/helpers"
	"eliest/models"
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

func (handler *EliestHandler) WinsCode(w http.ResponseWriter, r *http.Request) {

	var winPayload models.WinPayload

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
	hash := helpers.WinningHash(validator, serial)

	win, err := handler.Db.FindWinning(&hash)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		helpers.RespondWithError(w, http.StatusNotFound, "Invalid or used winning code")
		return
	}
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, GeneralServiceError)
		return
	}
	if win.Status != "used" && win.Status == "active" {
		err = handler.Db.UpdateWinningMap(win, map[string]interface{}{"status": "used"})
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, GeneralServiceError)
			return
		}
		helpers.RespondWithJSON(w, http.StatusOK, models.WinCodeCheckPayload{Amount: win.Amount, Status: "active"})
		return
	} else {
		helpers.RespondWithError(w, http.StatusBadRequest, "Your entered an invalid code")
		return
	}
}
func (handler *EliestHandler) DepositsFailed(w http.ResponseWriter, r *http.Request) {
	var winPayload models.WinPayload

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
	hash := helpers.WinningHash(validator, serial)

	win, err := handler.Db.FindWinning(&hash)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid winning code")
		return 
	}
	if win.Status == "used" {
		err = handler.Db.UpdateWinningMap(win, map[string]interface{}{"status": "active"})
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, GeneralServiceError)
			return
		}
		helpers.RespondWithJSON(w, http.StatusOK, "Winning code reversed")
		return
	} else {
		helpers.RespondWithError(w, http.StatusOK, "Invalid request")
		return
	}
}
func (handler *EliestHandler) DepositsSuccess(w http.ResponseWriter, r *http.Request) {

	var winPayload models.TransferredCallback

	err := json.NewDecoder(r.Body).Decode(&winPayload)
	defer r.Body.Close()
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, "")
}

func (handler *EliestHandler) TransferWinToAgent(w http.ResponseWriter, r *http.Request) {

	var winPayload models.TransferredCallback

	err := json.NewDecoder(r.Body).Decode(&winPayload)
	defer r.Body.Close()
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	_, err = findAgent(winPayload.Agent)
	if err != nil {
		helpers.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	if len(winPayload.Code) != 7 {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid winning code")
		return
	}
	a := []rune(winPayload.Code)
	pin := string(a[0:3])
	serial := string(a[3:7])
	validator := pin + serial
	hash := helpers.WinningHash(validator, serial)

	win, err := handler.Db.FindWinning(&hash)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid winning code")
		return 
	}
	if win.Status != "used" {
		err = handler.Db.UpdateWinningMap(win, map[string]interface{}{"status": "used"})
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, GeneralServiceError)
			return
		}
		//create transaction for agent
		helpers.RespondWithJSON(w, http.StatusOK, "You have successfully transferred your winning")
		return
	} else {
		helpers.RespondWithError(w, http.StatusBadRequest, "Invalid winning code")
	}

}

var agents = map[string]string{
	"08069475323": "Proud Nigerian",
	"08055913141": "Ade Femi",
	"08037122080": "Jacob femi",
}

func findAgent(phone string) (string, error) {
	pass, ok := agents[phone]
	if !ok {
		return "", errors.New("Agent Not found")
	}
	return pass, nil
}

//9509925
