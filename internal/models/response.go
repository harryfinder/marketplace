package models

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response ...
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}
type ResponseUser struct {
	Id    int64  `json:"id"`
	Token string `json:"token"`
	Role  int64  `json:"role"`
}

// Payload ...
type Payload struct {
	ID    int64    `json:"id"`
	Exp   int64    `json:"exp"`
	Roles []string `json:"roles"`
}

// Send ...
func (res *Response) Send(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)

	if res.Payload == nil && res.Code != http.StatusOK {
		res.Payload = struct {
			Error   bool   `json:"error,omitempty"`
			Message string `json:"message,omitempty"`
		}{
			Error:   true,
			Message: res.Message,
		}
	}
	if len(res.Message) == 0 {
		res.Message = http.StatusText(res.Code)
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("ERROR Sending response failed:", err)
	}
}
