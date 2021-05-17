package handler

import (
	"bytes"
	"eliest/helpers"
	"eliest/logger"
	"eliest/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/twinj/uuid"
)

var(logging logger.LogHandler
)


func (handler *EliestHandler)GetDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var res models.CoralDetailResponse
	res.TraceId = helpers.RandUpperAlpha(9)
	logging = logger.NewCoralZapLoger()
	username, password, ok := r.BasicAuth()
	bb, _ := ioutil.ReadAll(r.Body)
	body := fmt.Sprintf(" %s - %s - %s - %v", string(bb), r.URL.RequestURI(), r.Host, r.Header)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bb))
	defer r.Body.Close()
	if !ok {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		res.ResponseCode = "03"
		res.DisplayMessage = fmt.Sprintf("No basic auth present")
		res.CustomerName = "nil"

		logging.LogError(fmt.Sprintf("No basic auth present %s", body))

		details, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		w.Write(details)
		return
	}

	if !isAuthorised(username, password) {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		res.ResponseCode = "03"
		res.DisplayMessage = fmt.Sprintf("Invalid username or password")
		res.CustomerName = "nil"
		logging.LogError(fmt.Sprintf("Invalid username or password %s", body))
		details, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		w.Write(details)
		return
	}

	var req models.CoralDetailPayload

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		res.ResponseCode = "03"
		res.DisplayMessage = fmt.Sprintf("Err processing request - %v", err.Error())
		res.CustomerName = "nil"
		logging.LogError(fmt.Sprintf("Err processing request: %v - %v", err.Error(), body))

		details, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		w.Write(details)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.CustomerRef) == 0 || len(req.MerchantId) == 0 || req.MerchantId != "4032ELT10000001" {
		res.ResponseCode = "03"
		res.DisplayMessage = "Invalid Reqest payload"
		res.CustomerName = "nil"
		logging.LogError(fmt.Sprintf("Invalid Reqest payload - %v", body))

		w.WriteHeader(http.StatusBadRequest)

		details, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		w.Write(details)
		return
	}

	user, err := handler.Db.FindAccount(&models.Account{RefCode: req.CustomerRef})

	if user == nil{
		agent, err := handler.Db.FindAgent(&models.Agent{RefCode: req.CustomerRef})

		if err != nil {
			res.ResponseCode = "03"
			res.DisplayMessage = "Not Found - Invalid customer ref"
			res.CustomerName = "nil"
			logging.LogError(fmt.Sprintf("Not Found - Invalid customer ref - %v", body))
	
			w.WriteHeader(http.StatusBadRequest)
	
			details, err := json.Marshal(res)
			if err != nil {
				panic(err)
			}
	
			w.Write(details)
			return
		}

		
		res.ResponseCode = "00"
		res.DisplayMessage = fmt.Sprintf("%s Agent top up", agent.Phone)
		res.CustomerName = agent.Phone

		log.Printf("setting %v", res.TraceId)
		err = handler.RedisClient.Client.Set(res.TraceId, agent.Phone, time.Hour *2).Err()
		log.Println(err)
		
		logging.LogInfo(fmt.Sprintf("wallet top up - %v", body))

		w.WriteHeader(http.StatusOK)

		details, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		w.Write(details)
		return
	}

	if err != nil {
		res.ResponseCode = "03"
		res.DisplayMessage = "Not Found - Invalid customer ref"
		res.CustomerName = "nil"
		logging.LogError(fmt.Sprintf("Not Found - Invalid customer ref - %v", body))

		w.WriteHeader(http.StatusBadRequest)

		details, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		w.Write(details)
		return
	}

	res.ResponseCode = "00"
	res.DisplayMessage = fmt.Sprintf("%s wallet top up", user.MSISDN)
	res.CustomerName = user.MSISDN

	logging.LogInfo(fmt.Sprintf("wallet top up - %v", body))

	err = handler.RedisClient.Client.Set(res.TraceId, user.MSISDN, time.Minute *2).Err()
	log.Println(err)

	w.WriteHeader(http.StatusOK)

	details, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	w.Write(details)
}

func  (handler *EliestHandler)Notification(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var res models.CoralNotifResponse
	logging = logger.NewCoralZapLoger()
	bb, _ := ioutil.ReadAll(r.Body)
	body := fmt.Sprintf(" %s - %s - %s - %v", string(bb), r.URL.RequestURI(), r.Host, r.Header)
	defer r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bb))

	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		res.ResponseCode = "03"
		res.ResponseMessage = fmt.Sprintf("No basic auth present")
		logging.LogError(fmt.Sprintf("No basic auth present - %v", body))

		details, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		w.Write(details)
		return
	}

	if !isAuthorised(username, password) {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		res.ResponseCode = "03"
		res.ResponseMessage = fmt.Sprintf("Invalid username or password")
	logging.LogError(fmt.Sprintf("Invalid username or password - %v", body))

		details, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		w.Write(details)
		return
	}

	var req models.CoralNotifPayload

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		res.ResponseCode = "03"
		res.ResponseMessage = fmt.Sprintf("Err processing request - %v", err.Error())
		logging.LogError(fmt.Sprintf("Err processing request %v - %v", err.Error(), body))

		details, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		w.Write(details)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(req.CustomerRef) == 0 || len(req.MerchantId) == 0 || req.Amount < 0|| len(req.Hash) == 0 || len(req.TraceId) == 0 || req.MerchantId != "4032ELT10000001"{
		res.ResponseCode = "03"
		res.ResponseMessage = "Invalid Reqest payload"
		logging.LogError(fmt.Sprintf("Invalid Reqest payload - %v", body))

		w.WriteHeader(http.StatusBadRequest)

		details, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		w.Write(details)
		return
	}


	hash := req.HashValue()
	fmt.Printf("%s\n%s \n looks equal", hash, req.Hash)

	if hash != req.Hash {
		res.ResponseCode = "03"
		res.ResponseMessage = "Invalid hash sent"

		w.WriteHeader(http.StatusBadRequest)
		logging.LogError(fmt.Sprintf("Invalid hash sent - %v", body))
		details, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		w.Write(details)
		return
	}


log.Println(req.TraceId)


	ref := handler.RedisClient.Client.Get(req.TraceId).Val()
	log.Println(ref)
	if len(ref) == 10{
		_id := uuid.NewV4().String()
		wallet, _ := handler.Db.FindAgentWallet(&models.Wallet{Owner: req.CustomerRef, Class: "default"})
		if wallet != nil{
			_id = wallet.Id
		}
		handler.Db.CreateTransaction(&models.Transaction{Id: uuid.NewV4().String(), Amount: req.Amount, TRef: req.TraceId, Account: _id, Class: "credit", Status: "success"})
	}

	if len(ref) == 14{
		user, _ := handler.Db.FindAccount(&models.Account{RefCode: req.CustomerRef})
		handler.Db.UpdateUser(user, &models.Account{RefCode: req.CustomerRef, Balance: user.Balance + req.Amount})
	}

	res.ResponseCode = "00"
	res.ResponseMessage = fmt.Sprintf("Successful")
	logging.LogInfo(fmt.Sprintf("Successful - %v", body))

	w.WriteHeader(http.StatusOK)

	details, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	w.Write(details)
}

var users = map[string]string{
	"Bloomcan_Abingora": "WDsDdnJh___Vfdq86MF",
	"0test": "0secret",
	"d33Cana": "z#neh_WqHO^X",
}

func isAuthorised(username, password string) bool {
	pass, ok := users[username]
	if !ok {
		return false
	}
	return password == pass
}
