package main

import (
	"fmt"
	"log"
	"net/http"
	// "os"
	// "github.com/golang-jwt/jwt/v4"
)

var secret = []byte("SECRET_KEY")

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Token Verified")
}

func isAuthorized(endpoint func(http.Request, *http.ResponseWriter)) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func handleRequests() {
	http.Handle("/", isAuthorized(homepage))
	log.Fatal(http.ListenAndServe("9001", nil))
}

func main() {
	fmt.Println("Server Started")
	handleRequests()
}
