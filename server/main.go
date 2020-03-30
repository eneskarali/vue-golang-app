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

	log.Fatal(http.ListenAndServe(":8000", nil))
}
