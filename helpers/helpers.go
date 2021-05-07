package helpers

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/golang/gddo/httputil/header"
)



const intset = "0123456789"

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charset2 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

var seedRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandInt(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = intset[seedRand.Intn(len(intset))]
	}
	return string(b)
}

func RandUpperAlpha(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seedRand.Intn(len(charset))]
	}
	return string(b)
}

func RandAlpha(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset2[seedRand.Intn(len(charset))]
	}
	return string(b)
}


func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return errors.New(msg)
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	//dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		return err
	}

return err
}