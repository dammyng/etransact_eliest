package handler_test

import "testing"

func TestDetDetails_unknownUser_noAuth(t *testing.T) {
}

/**
import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetDetails_unknownUser_noAuth(t *testing.T) {

	clearTable()

	var jsonStr = []byte(`{"customerRef": "gegeg", "merchantId":"4hhh6" }`)
	req, _ := http.NewRequest("POST", "/coralpay/getdetails", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)
	checkResponseCode(t, http.StatusUnauthorized, response)
	var result map[string]string
	json.Unmarshal(response.Body.Bytes(), &result)
	require.Equal(t, "No basic auth present", result["displayMessage"])
	require.Equal(t, "nil", result["customerName"])
	require.Equal(t, "03", result["responseCode"])
}

func TestDetDetails_unknownUser_withAuth(t *testing.T) {
	clearTable()
	var jsonStr = []byte(`{"customerRef": "gegeg", "merchantId":"4hhh6" }`)
	req, _ := http.NewRequest("POST", "/coralpay/getdetails", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("username1", "password123")
	response := executeRequest(req)
	checkResponseCode(t, http.StatusUnauthorized, response)
	var result map[string]string
	json.Unmarshal(response.Body.Bytes(), &result)
	require.Equal(t, "Invalid username or password", result["displayMessage"])
	require.Equal(t, "nil", result["customerName"])
	require.Equal(t, "03", result["responseCode"])
}

func TestDetDetails_knownUser_withAuth(t *testing.T) {
	clearTable()
	acct := models.Account{
		MSISDN:  "08069475323",
		YOB:     "2001",
		Pin:     "2222",
		Balance: 0.00,
		RefCode: "543210",
		Referer: "",
		Status:  "active"}
	err := coralpay.DBConnection().CreateAccount(acct)
	if err != nil {
		log.Println(err)
	}
	var jsonStr = []byte(`{"customerRef": "543210", "merchantId":"1057ELL10000001"}`)
	req, _ := http.NewRequest("POST", "/coralpay/getdetails", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("0test", "0secret")
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response)
	var result map[string]string
	json.Unmarshal(response.Body.Bytes(), &result)
	require.Equal(t, fmt.Sprintf("%s wallet top up", acct.MSISDN), result["displayMessage"])
	require.Equal(t, acct.MSISDN, result["customerName"])
	require.Equal(t, "00", result["responseCode"])
}

func TestNotif_unknownUser_noAuth(t *testing.T) {
	clearTable()

	notif := `{
		"passBackReference": "12341820220", "traceId": "12341820220",
		"paymentReference": "03848982793092", "customerRef": "12345", "responseCode":"00", "merchantId":"93493MIU93832", "mobileNumber":"0808***7273",
		"amount":1000, "transactionDate":"2020-07-01T18:20:16.465107+01:00", "shortCode":"894",
		"currency":"NGN",
		"channel":"USSD",
		"hash": "093483ierl8w7s0s0-skj2j3k3wsj383mdlw838",
		}`
	var jsonStr = []byte(notif)
	req, _ := http.NewRequest("POST", "/coralpay/notification", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)
	var result map[string]string
	json.Unmarshal(response.Body.Bytes(), &result)
	fmt.Println(result)
	checkResponseCode(t, http.StatusUnauthorized, response)
	require.Equal(t, "No basic auth present", result["responseMessage"])
	require.Equal(t, "03", result["responseCode"])

}

func TestNotif_unknownUser_withAuth(t *testing.T) {
	clearTable()
	notif := `{
		"passBackReference": "12341820220", "traceId": "12341820220",
		"paymentReference": "03848982793092", "customerRef": "12345", "responseCode":"00", "merchantId":"93493MIU93832", "mobileNumber":"0808***7273",
		"amount":1000, "transactionDate":"2020-07-01T18:20:16.465107+01:00", "shortCode":"894",
		"currency":"NGN",
		"channel":"USSD",
		"hash": "093483ierl8w7s0s0-skj2j3k3wsj383mdlw838",
		}`
	var jsonStr = []byte(notif)
	req, _ := http.NewRequest("POST", "/coralpay/notification", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("username1", "password123")
	response := executeRequest(req)
	var result map[string]string
	json.Unmarshal(response.Body.Bytes(), &result)
	checkResponseCode(t, http.StatusUnauthorized, response)
	require.Equal(t, "Invalid username or password", result["responseMessage"])
	require.Equal(t, "03", result["responseCode"])
}

func TestNotification_knownUser(t *testing.T) {
	clearTable()

	notif := `{"paymentReference":"21032924010103335609",
	"customerRef":"5432104",
	"responseCode":"00",
	"merchantId":"1057ELL10000001",
	"mobileNumber":"0809***7503",
	"amount":3452.00,
	"passBackReference":"f7bb94a4-272c-4671-9074-d99d9704a9fc",
	"hash":"2cfac4d052bb30d183e6146ba29aa7c3e2d632db9b0e41c91fd1f0be3581972e",
	"transactionDate":"2021-03-29T10:33:29.8990613+01:00",
	"billerCode":"479000",
	"billerPrefix":"479",
	"traceId":"4HD5WSZ1I",
	"shortCode":"889",
	"currency":"NGN",
	"Channel":"USSD"}`
	var jsonStr = []byte(notif)
	req, _ := http.NewRequest("POST", "/coralpay/notification", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("0test", "0secret")
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response)
	var result map[string]string
	json.Unmarshal(response.Body.Bytes(), &result)
	require.Equal(t, "Successful", result["responseMessage"])
	require.Equal(t, "00", result["responseCode"])
}

var (
	db          = coralpay.DBConnection()
	gmloger     = gameslog.NewGamesFileSystem("", "")
	redisClient = myredis.NewRedisClient("", "")
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	appRouter := router.InitRoutes(db, gmloger, redisClient)
	appRouter.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected int, actual *httptest.ResponseRecorder) {
	if expected != actual.Code {
		t.Errorf("Expected response code %d. Got %d - with body %v\n", expected, actual.Code, actual.Body)
	}
}

func checkResponseCodeR(t *testing.T, expected int, actual *httptest.ResponseRecorder, to string) {
	require.Equal(t, []string{to}, actual.Header()["Location"])
	require.Equal(t, expected, actual.Code)
}

func clearTable() {
	db.Run("Delete FROM accounts")
	db.Run("Delete FROM agents")
	db.Run("Delete FROM transactions")
	db.Run("Delete FROM winnings")
}
**/