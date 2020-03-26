package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Set-Cookie")
}

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

func Signin(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var userCreds UserCredentials

	err := json.NewDecoder(r.Body).Decode(&userCreds)

	if err != nil {
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
			ExpiresAt: expirationTime.Unix(), // exipry time JTW de unix milisaniye cinsinden isteniyor
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //header ve payload ile token olusturuldu
	tokenString, err := token.SignedString(jwtKey)             // secret key eklenerek imzalandi

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //secret key imzası sirasinda hata olusursa
		return
	}

	http.SetCookie(w, &http.Cookie{ //jwt tokenimizi cookie olarak ekliyoruz, sonrasında client tarafında kontrol edebilmek icin
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func youValidUser(w http.ResponseWriter, r *http.Request) {

	cookieToken, err := r.Cookie("token")

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized) //eger herhangi bir cookie okunamazsa giris yapılmamıs demektir
			return
		}
		w.WriteHeader(http.StatusBadRequest) // farklı bir hata varsa yapılan istek ile ilgili bir problem var
		return
	}

	tokenText := cookieToken.Value

	claims := &userClaims{}

	token, err := jwt.ParseWithClaims(tokenText, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized) // token imzası hatali
			return
		}
		w.WriteHeader(http.StatusBadRequest) //istek hatali gönderilmis
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized) //gönderilen token valid degil
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username))) // tüm kontroller basarili sekilde tamamlanilirsa welcome mesajı gonder

}
