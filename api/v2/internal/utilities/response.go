package utilities

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Msg  string      `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (r *Response) Send(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(r)
}

func Send(w http.ResponseWriter, msg string, data interface{}, code int) {
	var r Response
	r.Msg = msg
	r.Data = data
	r.Send(w, code)
}
