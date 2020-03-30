package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/signin", signin)
	http.HandleFunc("/youValid", youValidUser)
	http.HandleFunc("/refresh", refreshToken)
	http.HandleFunc("/addpost", addPost)
	http.HandleFunc("/get10post", get10Post)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
