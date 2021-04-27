package router

import (
	"eliest/internals/db"
	"eliest/pkg/handler"

	"github.com/gorilla/mux"
)

func InitRoutes(db db.Handler) *mux.Router {
	var r = mux.NewRouter()

	handler := handler.NewEliestHandler(db)


	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/register", handler.Register).Methods("POST")
	v1.HandleFunc("/details/{msisdn}", handler.Details).Methods("GET")
	v1.HandleFunc("/fund", handler.Fund).Methods("POST")

	return r
}