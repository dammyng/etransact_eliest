package router

import (
	"eliest/internals/db"
	"eliest/logger/gamelogger"
	"eliest/myredis"
	"eliest/pkg/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes(db db.Handler, gameLogger gamelogger.GamesLogger, redis myredis.RedisClient) *mux.Router {
	var r = mux.NewRouter()

	handler := handler.NewEliestHandler(db, gameLogger, redis)


	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/register", handler.Register).Methods("POST")
	v1.HandleFunc("/details/{msisdn}", handler.Details).Methods("GET")
	v1.HandleFunc("/fund", handler.Fund).Methods("POST")
	v1.HandleFunc("/fund/voucher", handler.RechargeVoucher).Methods("POST")
	v1.HandleFunc("/play", handler.PlayGame).Methods("POST")
	v1.HandleFunc("/wins/validate", handler.WinsCode).Methods("POST")
	v1.HandleFunc("/wins/transfer_failed", handler.DepositsFailed).Methods("POST")
	v1.HandleFunc("/wins/transfer_success", handler.DepositsSuccess).Methods("POST")
	v1.HandleFunc("/wins/transfer_to_agent", handler.TransferWinToAgent).Methods("POST")
	v1.HandleFunc("/games/list", handler.GameList).Methods("GET")
	v1.HandleFunc("/agent/{id}", handler.FindAgent).Methods("GET")
	v1.HandleFunc("/winvalue/{code}", handler.WinValue).Methods("GET")

	coral := r.PathPrefix("/coralpay").Subrouter()
	coral.HandleFunc("/getdetails", handler.GetDetails).Methods("POST")
	coral.HandleFunc("/notification", handler.Notification).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("/var/www/eliest/static/")))
	r.HandleFunc("/", handler.Landing)

	return r
}