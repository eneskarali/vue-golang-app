package server

import (
	"log"
	"net/http"
)

func main()  {

	http.HandleFunc("/signin",Singin)
	http.HandleFunc("/dashboard",Dashboard)
	http.HandleFunc("/refresh",Refresh)

	log.Fatal(http.ListenAndServe(":8000",nil))
}