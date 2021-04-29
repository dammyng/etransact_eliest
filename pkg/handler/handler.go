package handler

import (
	"eliest/internals/db"
	"eliest/logger/gamelogger"
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
