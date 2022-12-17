package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Старт")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/user/registration", userRegistrationHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
