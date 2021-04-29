package handler

import (
	"eliest/helpers"
	"eliest/internals/data"
	"eliest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (handler *EliestHandler) PlayGame(w http.ResponseWriter, r *http.Request) {
	var reg models.GamePlayPayload


	err := json.NewDecoder(r.Body).Decode(&reg)
	defer r.Body.Close()
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	gameDetail, err := data.UseFindGame(reg.GameID)
	if err != nil {
		_ = "Invalid entry"
		helpers.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	entry := models.GameEntry{
		Phone:    reg.MSISDN,
		GuessOne: reg.Guess,
		GameId:   gameDetail.Id,
		Time:     time.Now(),
	}

	handler.GamesLogger.AddGameToCollection(gameDetail.Id, entry.String())
	currentCount, err := handler.GamesLogger.CollectionLength(gameDetail.Id)

	if int(gameDetail.TargetCount) == currentCount {
		handler.GamesLogger.ArchiveCollection(gameDetail.Id)
		amount := (gameDetail.WinnersCut / 100.00) * float64((gameDetail.TargetCount * gameDetail.Cost))
		winning, code := helpers.GererateWinning(amount, reg.MSISDN)
		handler.Db.CreateWinning(&winning)
		response := fmt.Sprintf(WinNote, code, amount)
		helpers.RespondWithJSON(w, http.StatusOK, response)
	}else {
		response := "Whoops! Sorry you didn’t win. Try again"
		helpers.RespondWithError(w, http.StatusNotFound, response)
		return
	}

}