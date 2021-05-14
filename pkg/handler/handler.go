package handler

import (
	"eliest/internals/db"
	"eliest/logger/gamelogger"
	"html/template"
	"net/http"
	"path"
)

type EliestHandler struct {
	Db db.Handler
	GamesLogger gamelogger.GamesLogger
}

func NewEliestHandler(db db.Handler, gamelogger gamelogger.GamesLogger) *EliestHandler {
	
	return &EliestHandler{
		Db: db,
		GamesLogger: gamelogger,
	}
}

func (handler *EliestHandler) Landing(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("static", "index.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}