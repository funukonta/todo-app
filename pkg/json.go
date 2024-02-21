package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

// DecodeJsonReq to decode from json in body request to object (struct/map)
func DecodeJsonReq(r *http.Request, v any) error {
	err := json.NewDecoder(r.Body).Decode(v)
	defer r.Body.Close()
	return err
}

func WriteJson(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Println(err)
	}
}

type MsgError struct {
	Error error `json:"error"`
}

type MsgOk struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func JsonErr(w http.ResponseWriter, statusCode int, err error) {
	WriteJson(w, statusCode, MsgError{Error: err})
}

func JsonOK(w http.ResponseWriter, statusCode int, message string, v any) {
	WriteJson(w, statusCode, MsgOk{Data: v, Message: message})
}
