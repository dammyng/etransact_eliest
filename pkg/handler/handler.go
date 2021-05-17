package handler

import (
	"eliest/internals/db"
	"eliest/logger/gamelogger"
	"eliest/myredis"
	"html/template"
	"net/http"
	"path"
)

type EliestHandler struct {
	Db db.Handler
	GamesLogger gamelogger.GamesLogger
	RedisClient myredis.RedisClient
}

func NewEliestHandler(db db.Handler, gamelogger gamelogger.GamesLogger, redis myredis.RedisClient) *EliestHandler {
	
	return &EliestHandler{
		Db: db,
		GamesLogger: gamelogger,
		RedisClient: redis,
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