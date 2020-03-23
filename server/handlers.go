package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)



var jwtKey = []byte("keep_life_simple")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type UserCredentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type userClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signin(w http.ResponseWriter, r *http.Request)  {

	var userCreds UserCredentials

	fmt.Printf("geldim")

	err := json.NewDecoder(r.Body).Decode(&userCreds)
	
	if err != nil{
		//json body hatali gelirse badrequest dondur
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[userCreds.Username]

	//kullanıcı adı bulunamadiginda veya sifre ile eslesmediginde unauthorized status dondur
	if !ok || expectedPassword != userCreds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &userClaims{
		Username: userCreds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),  // exipry time JTW de unix milisaniye cinsinden isteniyor
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)  //header ve payload ile token olusturuldu
	tokenString, err := token.SignedString(jwtKey) // secret key eklenerek imzalandi

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)  //secret key imzası sirasinda hata olusursa
		return
	}

	http.SetCookie(w, &http.Cookie{   //jwt tokenimizi cookie olarak ekliyoruz, sonrasında client tarafında kontrol edebilmek icin
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}