package pkg

import (
	"encoding/json"
	"net/http"
)

// DecodeJsonReq to decode from Request body
func DecodeJsonReq(r *http.Request, v any) error {
	err := json.NewDecoder(r.Body).Decode(v)
	defer r.Body.Close()
	return err
}

// WriteJsonRes to encode object (struct or map) to json format
func WriteJson(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(v)
}

type MsgSuccess struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
}

// WriteSuccessJson to write json on success
func WriteSuccessJson(w http.ResponseWriter, statuscode int, v any, pesan string) {
	WriteJson(w, statuscode, MsgSuccess{Data: v, Message: pesan})
}

type MsgError struct {
	Error error `json:"error"`
}

// WriteErrorJson to write json on Error
func WriteErrorJson(w http.ResponseWriter, statuscode int, err error) {
	WriteJson(w, statuscode, MsgError{Error: err})
}
