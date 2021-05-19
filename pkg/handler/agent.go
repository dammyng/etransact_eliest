package handler

import (
	"eliest/helpers"
	"net/http"

	"github.com/gorilla/mux"
)


func (handler *EliestHandler) FindAgent(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	agentId := params["id"]

	agent, err := findAgent(agentId)
	if err != nil {
		helpers.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, agent)
}