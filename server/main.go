package main

import (
	"log"
	"fmt"
	"net/http"
)

func main()  {

	http.HandleFunc("/signin",Signin)
	//http.HandleFunc("/dashboard",Dashboard)
	//http.HandleFunc("/refresh",Refresh)
	fmt.Printf("dinliyorum kral")
	log.Fatal(http.ListenAndServe(":8000",nil))
}