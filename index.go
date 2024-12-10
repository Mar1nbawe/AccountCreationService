package main

import (
	Endpoints "AccountCreationService/Endpoints"
	"log"
	"net/http"
)

func main() {

	//var password string
	//fmt.Printf("Please enter your password: ")
	//fmt.Scanln(&password)
	//
	//start := time.Now()
	//hash, _ := encrypt.HashPassword(password)
	//
	//fmt.Println("password:", password)
	//fmt.Println("hash:", hash)
	//
	//match := encrypt.VerifyPassword(password, hash)
	//duration := time.Since(start)
	//fmt.Println("match: ", match, " in :", duration)

	http.HandleFunc("/user", Endpoints.Userhandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))

}
