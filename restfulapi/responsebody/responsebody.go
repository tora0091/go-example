package responsebody

import (
	"encoding/json"
	"net/http"

	"go-example/restfulapi/model"
)

type HttpStatus struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type HttpStatusWithUsers struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    model.Users `json:"data"`
}

type HttpStatusWithUser struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    model.User `json:"data"`
}

func StatusOK(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HttpStatus{Status: http.StatusOK, Message: http.StatusText(http.StatusOK)})
}

func StatusBadRequest(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	if msg == "" {
		msg = http.StatusText(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(HttpStatus{Status: http.StatusBadRequest, Message: msg})
}

func StatusInternalServerError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	if msg == "" {
		msg = http.StatusText(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(HttpStatus{Status: http.StatusInternalServerError, Message: msg})
}

func StatusOKWithUsers(w http.ResponseWriter, data model.Users) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HttpStatusWithUsers{Status: http.StatusOK, Message: http.StatusText(http.StatusOK), Data: data})
}

func StatusOKWithUser(w http.ResponseWriter, data model.User) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HttpStatusWithUser{Status: http.StatusOK, Message: http.StatusText(http.StatusOK), Data: data})
}
