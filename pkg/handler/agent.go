package handler

import (
	"eliest/helpers"
	"net/http"
	"unicode/utf8"

	"github.com/gorilla/mux"
)


func (handler *EliestHandler) FindAgent(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	agentId := params["id"]

	agent, err := findAgent(trimFirstRune(agentId))
	if err != nil {
		helpers.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, agent)
}

func trimFirstRune(s string) string {
    _, i := utf8.DecodeRuneInString(s)
    return s[i:]
}