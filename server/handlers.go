package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var jwtKey = []byte("secretKey")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type UserCredentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func Singin(w http.ResponseWriter, r *http.Request)  {

	var userCreds UserCredentials

	err := json.NewDecoder(r.Body).Decode(&userCreds)
	
	if err != nil{
		//return bad request status when json body is wrong
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[userCreds.Username]

	//return unauthorized status when username not found or not match the password
	if !ok || expectedPassword != userCreds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}