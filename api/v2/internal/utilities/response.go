package utilities

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Msg     string      `json:"message,omitempty"`
	Payload interface{} `json:"data,omitempty"`
}

func (r *Response) Send(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

  if r.Msg == "" && r.Payload == nil {
		json.NewEncoder(w).Encode("Backend failure, Message and Payload are empty")
  } else if r.Msg == "" {
		json.NewEncoder(w).Encode(r.Payload)
	} else if r.Payload == nil {
		json.NewEncoder(w).Encode(r.Msg)
  } else {
		json.NewEncoder(w).Encode(Response{Msg: r.Msg, Payload: r.Payload})
	}
}

func Send(w http.ResponseWriter, msg string, data interface{}, code int) {
	var r Response
	r.Msg = msg
	r.Payload = data
	r.Send(w, code)
}
