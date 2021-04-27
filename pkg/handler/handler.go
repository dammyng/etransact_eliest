package handler

import "eliest/internals/db"

type EliestHandler struct {
	Db db.Handler
}

func NewEliestHandler(db db.Handler) *EliestHandler {
	
	return &EliestHandler{
		Db: db,
	}
}
