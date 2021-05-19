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

type Game struct {
	Id   string `json:"id"`
	Text string `json:"test"`
}

func (handler *EliestHandler) GameList(w http.ResponseWriter, r *http.Request) {

	var games []Game
	one := Game{
		Id:   "1",
		Text: "Tanzanite #50 win #4000",
	}
	two := Game{
		Id:   "2",
		Text: "Alexandrite #100 win #8,000",
	}
	three := Game{
		Id:   "3",
		Text: "Jadeite #200 win #16,000",
	}
	four := Game{
		Id:   "4",
		Text: " Pink Diamond #500 win #40,000",
	}
	games = append(games, one)
	games = append(games, two)
	games = append(games, three)
	games = append(games, four)

	helpers.RespondWithJSON(w, http.StatusOK, games)
}

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
	
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, GeneralServiceError)
		return
	}

	user, err := handler.Db.FindAccount(&models.Account{MSISDN: reg.MSISDN})
	if err != nil {
		helpers.RespondWithError(w, http.StatusNotFound, UserNotFound)
		return
	}

	err = handler.Db.UpdateUser(user, &models.Account{Balance: user.Balance - float64(gameDetail.Cost)})
	
	if err != nil {
		helpers.RespondWithError(w, http.StatusNotFound, UserNotFound)
		return
	}

	if int(gameDetail.TargetCount) == currentCount {
		handler.GamesLogger.ArchiveCollection(gameDetail.Id)
		amount := (gameDetail.WinnersCut / 100.00) * float64((gameDetail.TargetCount * gameDetail.Cost))
		winning, code := helpers.GererateWinning(amount, reg.MSISDN)
		handler.Db.CreateWinning(&winning)
		response := fmt.Sprintf(WinNote, code, amount)
		helpers.RespondWithJSON(w, http.StatusOK, response)
	} else {
		response := "Whoops! Sorry you didn’t win. Try again"
		helpers.RespondWithError(w, http.StatusNotFound, response)
		return
	}

}
