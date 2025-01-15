package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/refresh-token", refreshTokenHandler)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("Error starting server: ", err.Error())
	}
}
